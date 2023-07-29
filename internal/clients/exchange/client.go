package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/EclipIIse/exchanger/internal/models"
	"github.com/pkg/errors"
)

type Client struct {
	host string
}

func New(host string) *Client {
	return &Client{host: host}
}

func (c *Client) request(method string, data []byte) (int, []byte, error) {
	const timeout = time.Second * 10

	requestBody := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, c.host+method, requestBody)
	if err != nil {
		return 0, nil, errors.Wrap(err, "exchange create request fail")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: timeout}
	response, err := client.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err, "exchange request failed")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "read exchange response body failed")
	}

	return response.StatusCode, responseBody, nil
}

func (c *Client) GetCurrency(request *models.MyModel) (*models.MyModel, error) {
	const method = ""

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request failed")
	}

	code, responseBody, err := c.request(method, requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "sales funnel request failed")
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("%s retrun code %d with body %s", method, code, string(responseBody))
	}

	response := new(models.MyModel)
	if err = json.Unmarshal(responseBody, response); err != nil {
		return nil, errors.Wrap(err, "currency unmarshall response failed")
	}

	return response, nil
}
