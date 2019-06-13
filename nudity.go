package deepAI

import (
	"bytes"
	"mime/multipart"
	"os"

	"deepAI/models"
)

func (c *Client) DetectWithURL(url string) (models.NudityResponse, error) {
	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)
	defer w.Close()

	fw, err := w.CreateFormField("image")
	if err != nil {
		return models.NudityResponse{}, err
	}

	_, err = fw.Write([]byte(url))
	if err != nil {
		return models.NudityResponse{}, err
	}

	nudityResponse, err := c.request(buf, w)
	if err != nil {
		return models.NudityResponse{}, err
	}

	return nudityResponse, nil
}

func (c *Client) DetectWithFile(data []byte, fileName string) (models.NudityResponse, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open(fileName)
	if err != nil {
		return models.NudityResponse{}, err
	}
	defer f.Close()

	fw, err := w.CreateFormFile("image", fileName)
	if err != nil {
		return models.NudityResponse{}, err
	}

	_, err = fw.Write(data)
	if err != nil {
		return models.NudityResponse{}, err
	}

	if _, err = fw.Write(data); err != nil {
		return models.NudityResponse{}, err
	}
	w.Close()

	nudityResponse, err := c.request(b, w)
	if err != nil {
		return models.NudityResponse{}, err
	}

	return nudityResponse, nil
}
