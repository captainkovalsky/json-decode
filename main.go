package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.PanicLevel)
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		"http://demo1770324.mockable.io/me",
		nil)

	if err != nil {
		log.WithField("error", err).Error("Unable to create POST request")
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.WithField("error", err).Error("Unable to execute request")
		return
	}

	var response struct {
		Msg string `json:"msg"`
	}

	errParse := json.NewDecoder(res.Body).Decode(&response)

	if errParse != nil {
		log.WithField("error", err).Error("Unable to parse the response from Request")
		return
	}

	log.WithField("result", response).Info("Parsed Response")
}
