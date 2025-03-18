package IntermediateServices

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PubSubMessage struct {
	Message struct {
		Data      string `json:"data"`
		MessageID string `json:"message_id"`
	} `json:"message"`
}

func jsonDecoderWithNumber(data []byte) (map[string]interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.UseNumber()
	var result map[string]interface{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func PubsubReqValidation(c *fiber.Ctx) (map[string]interface{}, string, error) {
	fmt.Println("c.body ::", c.Body())
	if len(c.Body()) == 0 {
		fmt.Println("ERROR: No message received in the body.")
		return nil, "", fmt.Errorf("no message received in the body")
	}

	var pubSubMessage PubSubMessage
	if err := c.BodyParser(&pubSubMessage); err != nil {
		fmt.Println("ERROR: Invalid Pub/Sub message format.")
		return nil, "", fmt.Errorf("invalid Pub/Sub message format")
	}
	fmt.Println("pubsubmessage after bodyparser  ::", pubSubMessage)
	fmt.Println("pubsubmessage after bodyparser  ::", pubSubMessage.Message)
	fmt.Println("pubsubmessage after bodyparser  ::", pubSubMessage.Message.Data)

	if pubSubMessage.Message.Data == "" {
		fmt.Println("ERROR: 'data' field is missing or empty.")
		return nil, "", fmt.Errorf("'data' field is missing or empty")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(pubSubMessage.Message.Data)
	if err != nil {
		fmt.Println("ERROR: Failed to decode Base64 data.")
		return nil, "", fmt.Errorf("failed to decode Base64 data")
	}

	

	decodedData, err := jsonDecoderWithNumber(decodedBytes)
	if err != nil {
		fmt.Println("ERROR: Decoded data is not valid JSON.")
		return nil, "", fmt.Errorf("decoded data is not valid JSON")
	}

	

	fmt.Println("Data after unmarshalling:", decodedData)
	return decodedData, pubSubMessage.Message.MessageID, nil
}
