package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PartialBeaconBellatrixBlock struct {
	Body struct {
		ExecutionPayload struct {
			BlockHash string `json:"block_hash"`
		} `json:"execution_payload"`
	} `json:"body"`
}

type EthClient struct {
	baseURL string
}

func NewEthClient(baseURL string) *EthClient {
	return &EthClient{baseURL: baseURL}
}

func (c *EthClient) GetPartialBeaconBellatrixBlock(slot uint64) (*PartialBeaconBellatrixBlock, error) {
	var res struct {
		Data struct {
			Message PartialBeaconBellatrixBlock `json:"message"`
		} `json:"data"`
	}

	resp, err := http.Get(fmt.Sprintf("%s/eth/v2/beacon/blocks/%d", c.baseURL, slot))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res.Data.Message, resp.Body.Close()
}
