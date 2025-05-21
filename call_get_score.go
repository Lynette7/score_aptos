package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
)

func main() {
	accountAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"
	userAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"

	functionID := fmt.Sprintf("%s::get_score::get_score", accountAddress)

	url := "https://api.testnet.aptoslabs.com/v1/view"

	payload := []byte(`{
		"function": "` + functionID + `",
		"type_arguments": [],
		"arguments": ["` + userAddress + `"]
	}`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("%s", body)

	var result []string
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return
	}

	if len(result) == 0 {
		fmt.Println("No result returned from the view function")
		return
	}

	score, ok := new(big.Int).SetString(result[0], 10)
	if !ok {
		fmt.Printf("Error converting score to big.Int: %s\n", result[0])
		return
	}

	fmt.Printf("Score: %s\n", score.String())
}
