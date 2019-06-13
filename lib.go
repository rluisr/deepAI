package deepAI

import (
	"bytes"
	"deepAI/models"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
)

func (c *Client) request(buf bytes.Buffer, w *multipart.Writer) (models.NudityResponse, error) {
	req, err := http.NewRequest(http.MethodPost, NudityURL, &buf)
	if err != nil {
		return models.NudityResponse{}, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("api-key", c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.NudityResponse{}, err
	}
	defer resp.Body.Close()

	buf = bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return models.NudityResponse{}, err
	}

	if resp.StatusCode != 200 {
		return models.NudityResponse{}, errors.New(buf.String())
	}

	var nudityResponse models.NudityResponse
	err = json.Unmarshal(buf.Bytes(), &nudityResponse)
	if err != nil {
		return models.NudityResponse{}, err
	}
	return nudityResponse, nil
}
