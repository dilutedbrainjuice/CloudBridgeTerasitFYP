const firebaseConfig = {
  apiKey: "AIzaSyAPLzQmeLSrda6NLvV32pojAdE3L3XbQL0",
  authDomain: "mychatapp-6e730.firebaseapp.com",
  projectId: "mychatapp-6e730",
  storageBucket: "mychatapp-6e730.appspot.com",
  messagingSenderId: "836162645657",
  appId: "1:836162645657:web:e716a1670c989e6b8fd208",
  measurementId: "G-W5F0HWLVQL"
};

// Initialize Firebase
firebase.initializeApp(firebaseConfig);
const db = firebase.firestore();

const messageContainer = document.getElementById('message-container');
const messageInput = document.getElementById('message-input');
const sendButton = document.getElementById('send-button');

// Function to add a message to the Firestore database
const sendMessage = (messageText) => {
  db.collection('messages').add({
    text: messageText,
    timestamp: firebase.firestore.FieldValue.serverTimestamp()
  });
};

// Function to display messages in the chat container
const displayMessages = (messages) => {
  messageContainer.innerHTML = '';
  messages.forEach((message) => {
    const messageDiv = document.createElement('div');
    messageDiv.textContent = message.text;
    messageContainer.appendChild(messageDiv);
  });
};

// Fetch messages from Firestore and display them
db.collection('messages').orderBy('timestamp').onSnapshot((snapshot) => {
  const messages = [];
  snapshot.forEach((doc) => {
    messages.push(doc.data());
  });
  displayMessages(messages);
});

// Event listener for sending messages
sendButton.addEventListener('click', () => {
  const messageText = messageInput.value.trim();
  if (messageText !== '') {
    sendMessage(messageText);
    messageInput.value = '';
  }
});