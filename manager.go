package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	/**
	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
	*/
	websocketUpgrader = websocket.Upgrader{
		// Apply the Origin Checker
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// checkOrigin will check origin and return true if its allowed
func checkOrigin(r *http.Request) bool {

	// Grab the request origin
	origin := r.Header.Get("Origin")

	switch origin {
	// Update this to HTTPS
	case "http://localhost:9000":
		return true
	default:
		return false
	}
}

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	clients ClientList
	db      *sql.DB

	// Using a syncMutex here to be able to lcok state before editing clients
	// Could also use Channels to block
	sync.RWMutex
	// handlers are functions that are used to handle Events
	handlers map[string]EventHandler
	// otps is a map of allowed OTP to accept connections from

}

// NewManager is used to initalize all the values inside the manager
func NewManager(ctx context.Context, db *sql.DB) *Manager {
	m := &Manager{
		clients:  make(ClientList),
		db:       db,
		handlers: make(map[string]EventHandler),
		// Create a new retentionMap that removes Otps older than 5 seconds

	}
	m.setupEventHandlers(db)
	return m
}

// setupEventHandlers configures and adds all handlers
func (m *Manager) setupEventHandlers(db *sql.DB) {
	m.handlers[EventSendMessage] = func(event Event, c *Client) error {
		return SendMessageHandler(db, event, c)
	}
	m.handlers[EventChangeRoom] = ChatRoomHandler
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (m *Manager) routeEvent(event Event, c *Client) error {
	// Check if Handler is present in Map
	if handler, ok := m.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

// serveWS is a HTTP Handler that the has the Manager that allows connections
func (m *Manager) serveWS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("New connection")
		// Begin by upgrading the HTTP request
		conn, err := websocketUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		// Create New Client
		client := NewClient(conn, m)
		// Add the newly created client to the manager
		m.addClient(client)

		go client.readMessages()
		go client.writeMessages()
	}
}

// addClient will add clients to our clientList
func (m *Manager) addClient(client *Client) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.clients[client] = true
}

// removeClient will remove the client and clean up
func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		// close connection
		client.connection.Close()
		// remove
		delete(m.clients, client)
	}
}
