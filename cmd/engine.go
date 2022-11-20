package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var origin string = "http://localhost:33566"

func sendRequestToEngine(method string, path string, data map[string]interface{}) map[string]interface{} {
	url := origin + path
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(data)
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	if len(body) == 0 {
		return nil
	} else if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		// nothing to do
	} else {
		log.Fatal(resp.StatusCode, resp.Status, result)
	}
	return result
}

func EngineGet(path string) map[string]interface{} {
	return sendRequestToEngine(http.MethodGet, path, nil)
}

func EnginePost(path string, data map[string]interface{}) map[string]interface{} {
	return sendRequestToEngine(http.MethodPost, path, data)
}

func EnginePut(path string, data map[string]interface{}) map[string]interface{} {
	return sendRequestToEngine(http.MethodPut, path, data)
}

func EnginePatch(path string, data map[string]interface{}) map[string]interface{} {
	return sendRequestToEngine(http.MethodPatch, path, data)
}

func EngineDelete(path string, data map[string]interface{}) map[string]interface{} {
	return sendRequestToEngine(http.MethodDelete, path, data)
}
