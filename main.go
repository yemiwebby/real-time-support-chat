package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-support-chat/chat"
	"real-time-support-chat/config"
)


func main() {
    config.LoadEnv()

    mux := http.NewServeMux()

    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/create-conversation", createConversationHandler)
    mux.HandleFunc("/add-participant", addParticipantHandler)
    http.HandleFunc("/send-message", sendMessageHandler)

    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Real-Time Support Chat System!")
}

func createConversationHandler(w http.ResponseWriter, r *http.Request) {
    conversationName := r.URL.Query().Get("name")
    if conversationName == "" {
        http.Error(w, "Conversation name is required", http.StatusBadRequest)
        return
    }

    conversationSid, err := chat.CreateConversation(conversationName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Conversation created successfully: %s\n", conversationSid)
}

func addParticipantHandler(w http.ResponseWriter, r *http.Request) {
    conversationSid := r.URL.Query().Get("conversationSid")
    identity := r.URL.Query().Get("identity")
    if conversationSid == "" || identity == "" {
        http.Error(w, "conversationSid and identity are required", http.StatusBadRequest)
        return
    }

    if err := chat.AddParticipant(conversationSid, identity); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Participant added successfully")
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
    conversationSid := r.URL.Query().Get("conversationSid")
    author := r.URL.Query().Get("author")
    body := r.URL.Query().Get("body")
    if conversationSid == "" || author == "" || body == "" {
        http.Error(w, "conversationSid, author, and body are required", http.StatusBadRequest)
        return
    }

    if err := chat.SendMessage(conversationSid, author, body); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, "Message sent successfully")

}