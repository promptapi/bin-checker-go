package binchecker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// Result represents incoming JSON data from Prompt API - BIN Checker
type Result struct {
	BankName string `json:"bank_name"`
	Country  string `json:"country"`
	URL      string `json:"url"`
	Type     string `json:"type"`
	Scheme   string `json:"scheme"`
	Bin      string `json:"bin"`
}

// ErrorResponse represents response errors from Prompt API - BIN Checker
type ErrorResponse struct {
	Message string `json:"message"`
}

var promptAPIEndpoint = "https://api.promptapi.com/bincheck/"

// BinChecker collects bin information from Prompt API - BIN Checker
func BinChecker(binNumber string, result *Result) error {
	apiKey, ok := os.LookupEnv("PROMPTAPI_TOKEN")
	if !ok {
		return errors.New("You need to set PROMPTAPI_TOKEN environment variable")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", promptAPIEndpoint+binNumber, nil)
	req.Header.Set("apikey", apiKey)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		msg := new(ErrorResponse)
		err := json.Unmarshal(body, msg)
		if err != nil {
			return err
		}
		return errors.New(msg.Message)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}
