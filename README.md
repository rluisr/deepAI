DeepAI Golang SDK
==================
now support only nudity API.

Usage
-----
```go
func main() {
	c := deepAI.Client{
		Token: "1111-1111-1111-1111",
	}

	// with URL
	nudityResponse, err := c.DetectWithURL("https://files-uploader.xzy.pw/test.jpeg")
	if err != nil {
		panic(err)
	}

	// with File
	data, err := ioutil.ReadFile("test.jpeg")
	if err != nil {
		panic(err)
	}
	nudityResponse, err = c.DetectWithFile(data, "test.jpeg")
	if err != nil {
		panic(err)
	}
}
```
