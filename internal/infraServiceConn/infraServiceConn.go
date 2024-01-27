package infraserviceconn

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SendFileToInfraService(path string) error {
	infrastructureURL := os.Getenv("INFRASTRUCTURE_SERVICE_ENDPOINT")

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reqBody, contentType, err := createMultipartBody(file, path)
	if err != nil {
		return err
	}

	return sendPOST(infrastructureURL, contentType, reqBody)
}

func createMultipartBody(file *os.File, path string) (*bytes.Buffer, string, error) {
	var reqBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&reqBody)

	filePart, err := multipartWriter.CreateFormFile("file", path)
	if err != nil {
		return nil, "", err
	}

	_, err = io.Copy(filePart, file)
	if err != nil {
		return nil, "", err
	}

	multipartWriter.Close()

	return &reqBody, multipartWriter.FormDataContentType(), nil
}

func sendPOST(url, contentType string, reqBody *bytes.Buffer) error {
	res, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
