package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Chat(msg string) {
		
	const url = "http://localhost:11434/api/chat"
	const historyFile = "history.json"
	
	type Message struct {
		Role string `json:"role"`
		Content string `json:"content"`
	}

	type ChatRequest struct {
		Model string `json:"model"`
		Messages []Message `json:"messages"`
		Stream bool `json:"stream"`
	}

	type ChatResponse struct {
		Model string `json:"model"`
		Created_at string `json:"created_at"`
		Message Message `json:"message"`
		Done bool `json:"done"`
	}

	var history []Message

	f,err := os.Open(historyFile)
	if err == nil{
		defer f.Close()
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&history)
	} else {
		history = []Message{}
	}

	newMessage := Message{
		Role: "user",
		Content: msg,
	}

	history = append(history,newMessage)

	payload := ChatRequest{
		Model: "deepseek-r1",
		Messages: history,
		Stream: true,
	}
	
	body, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	
	resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	var chatResp ChatResponse
	var reply string

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		err = json.Unmarshal(scanner.Bytes(),&chatResp)
		if err != nil{
			continue
		}
		fmt.Print(chatResp.Message.Content)
		reply+=chatResp.Message.Content
	}

	history = append(history, Message{
		Role: "assistant",
		Content: reply,
	})

	f, err = os.Create(historyFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(history)
}
