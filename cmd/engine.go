package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var origin string = "http://localhost:33566"

func sendRequestToEngine(method string, path string, data map[string]interface{}) (map[string]interface{}, int) {
	url := origin + path
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(data)
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
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
	} else if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	return result, resp.StatusCode
}

func EngineGet(path string) map[string]interface{} {
	res, _ := sendRequestToEngine(http.MethodGet, path, nil)
	return res
}

func EnginePost(path string, data map[string]interface{}) map[string]interface{} {
	res, _ := sendRequestToEngine(http.MethodPost, path, data)
	return res
}

func EnginePut(path string, data map[string]interface{}) map[string]interface{} {
	res, _ := sendRequestToEngine(http.MethodPut, path, data)
	return res
}

func EnginePatch(path string, data map[string]interface{}) map[string]interface{} {
	res, _ := sendRequestToEngine(http.MethodPatch, path, data)
	return res
}

func EngineDelete(path string, data map[string]interface{}) map[string]interface{} {
	res, _ := sendRequestToEngine(http.MethodDelete, path, data)
	return res
}
