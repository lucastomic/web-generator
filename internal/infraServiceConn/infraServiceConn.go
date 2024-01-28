package infraserviceconn

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/lucastomic/web-generator-service/internal/infraServiceConn/compression"
)

func SendFilesToInfraService(paths []string) error {
	infrastructureURL := os.Getenv("INFRASTRUCTURE_SERVICE_ENDPOINT")
	compressedPath := fmt.Sprintf("/tmp/%d.zip", rand.Int())
	_, err := compression.CompressFiles(compressedPath, paths)
	if err != nil {
		return err
	}
	file, err := os.Open(compressedPath)
	if err != nil {
		return err
	}
	defer os.Remove(compressedPath)

	reqBody, contentType, err := createMultipartBody(file, compressedPath)
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
