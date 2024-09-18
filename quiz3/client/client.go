package client

import (
	"bandProtocol/quiz3/constants"
	"bandProtocol/quiz3/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Broadcast  string
	Monitor    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		Broadcast:  "https://mock-node-wgqbnxruha-as.a.run.app/broadcast",
		Monitor:    "https://mock-node-wgqbnxruha-as.a.run.app/check/",
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) BroadcastTransaction(req *models.BroadcastTransactionRequest) (*models.BroadcastTransactionResponse, error) {
	// marshal req to JSON
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// create a new http.Request with the JSON payload
	httpReq, err := http.NewRequest(http.MethodPost, c.Broadcast, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// set the Content-Type header
	httpReq.Header.Set("Content-Type", "application/json")

	// send the request
	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// check status code
	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	// unmarshal the response body to a BroadcastTransactionResponse
	var resp models.BroadcastTransactionResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) MonitorTransaction(txHash string, interval time.Duration) {
	for {
		status, err := c.CheckTransactionStatus(txHash)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(status)

		if status == constants.CONFIRMED || status == constants.FAILED || status == constants.DNE {
			return
		}

		time.Sleep(interval)
	}
}

func (c *Client) CheckTransactionStatus(txHash string) (string, error) {
	monitorURL := fmt.Sprintf("%s%s", c.Monitor, txHash)

	// create a new http.Request
	httpReq, err := http.NewRequest(http.MethodGet, monitorURL, nil)
	if err != nil {
		return "", err
	}

	// send the request
	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer httpResp.Body.Close()

	// unmarshal the response body to a MonitorTransactionResponse
	var resp models.MonitorTransactionResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return "", err
	}

	return resp.TXStatus, nil
}
