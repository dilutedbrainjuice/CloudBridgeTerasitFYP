package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := ConnectPostgresDB()
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/home/", homehandler)
	http.HandleFunc("/registersite/", registersitehandler)
	http.HandleFunc("/registerform/", RegisterHandler(db))
	http.HandleFunc("/about/", abouthandler)
	http.HandleFunc("/userdashboard/", validateToken(dashboardshow))
	http.HandleFunc("/api/users", dashboardHandler(db))
	http.Handle("/users/images", http.FileServer(http.Dir("./uploads")))
	//Login page and login logic
	http.HandleFunc("/login/", loginhandler)
	http.HandleFunc("/loginform/", loginformhandler(db))
	http.HandleFunc("/logout/", logoutHandler)

	imageDir := "../CloudBridgeTerasitFYP/uploads/"

	// Create a file server handler to serve files from the image directory
	fs := http.FileServer(http.Dir(imageDir))

	// Handle requests for images by serving files from the file server
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	//initializing server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//display chat page
	http.HandleFunc("/initiateprivatechat", validateToken(chatHandler))
	//http.HandleFunc("/initiateprivatechat", validateToken(initiatePrivateChatHandler))
	http.HandleFunc("/api/currentuserIDName", validateToken(currentUserIDNameHandler))
	http.HandleFunc("/fetch-chat-history", FetchChatHistoryHandler(db))

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()
	manager := NewManager(ctx, db)
	http.HandleFunc("/ws", manager.serveWS())

	log.Println("Listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}
