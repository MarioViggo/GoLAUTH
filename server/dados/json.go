package dados

import (
	"encoding/json"
	"log"
)

func JSONMessage(message string) []byte {
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResp
}

func JSONRaw(message string) []byte {
	jsonResp, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResp
}
