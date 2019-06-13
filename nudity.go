package deepAI

import (
	"bytes"
	"mime/multipart"
)

func (c *Client) DetectWithURL(url string) (NudityResponse, error) {
	var buf bytes.Buffer

	w := multipart.NewWriter(&buf)
	defer w.Close()

	fw, err := w.CreateFormField("image")
	if err != nil {
		return NudityResponse{}, err
	}

	_, err = fw.Write([]byte(url))
	if err != nil {
		return NudityResponse{}, err
	}

	nudityResponse, err := c.request(buf, w)
	if err != nil {
		return NudityResponse{}, err
	}

	return nudityResponse, nil
}

func (c *Client) DetectWithFile(data []byte) (NudityResponse, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("image", "foo")
	if err != nil {
		return NudityResponse{}, err
	}

	_, err = fw.Write(data)
	if err != nil {
		return NudityResponse{}, err
	}

	if _, err = fw.Write(data); err != nil {
		return NudityResponse{}, err
	}
	w.Close()

	nudityResponse, err := c.request(b, w)
	if err != nil {
		return NudityResponse{}, err
	}

	return nudityResponse, nil
}
