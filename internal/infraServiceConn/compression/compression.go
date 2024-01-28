package compression

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"sync"
)

var (
	maxGoroutines      = 50
	goroutineSemaphore = make(chan struct{}, maxGoroutines)
)

func CompressFiles(outName string, paths []string) (*os.File, error) {
	outFile, err := os.Create(outName)
	if err != nil {
		return &os.File{}, nil
	}
	defer outFile.Close()
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()
	addFilesToZip(paths, zipWriter)
	return outFile, nil
}

func addFilesToZip(paths []string, zipWriter *zip.Writer) error {
	wg := sync.WaitGroup{}
	errChan := make(chan error, len(paths))
	for _, path := range paths {
		goroutineSemaphore <- struct{}{} // Acquire
		go addFileToZipConcurrent(path, zipWriter, errChan, &wg)
	}
	wg.Wait()
	close(errChan)
	for err := range errChan {
		return err
	}
	return nil
}

func addFileToZipConcurrent(
	path string,
	zipWriter *zip.Writer,
	errChan chan error,
	wg *sync.WaitGroup,
) {
	wg.Add(1)
	if err := addFileToZip(path, zipWriter); err != nil {
		errChan <- err
	}
	defer wg.Done()
	defer func() { <-goroutineSemaphore }() // Release
}

func addFileToZip(path string, zipWriter *zip.Writer) error {
	fileToAdd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileToAdd.Close()
	info, err := fileToAdd.Stat()
	if err != nil {
		return err
	}
	writer, err := createHeader(path, zipWriter, info)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToAdd)
	return err
}

func createHeader(path string, zipWriter *zip.Writer, info fs.FileInfo) (io.Writer, error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}
	header.Name = path
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return nil, err
	}
	return writer, nil
}
