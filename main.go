package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/JunNishimura/chatgpt-test/openai"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	requestBody := openai.Request{
		Model: "gpt-3.5-turbo",
		Messages: []*openai.Message{
			{
				Role:    "user",
				Content: "Say this is a test",
			},
		},
		Temperature: 0.7,
	}
	fmt.Printf("%+v\n", requestBody)

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", openai.URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response openai.Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Choices[0].Message.Content)
}
