package compression

import (
	"os"
	"testing"
	"time"

	"github.com/lucastomic/web-generator-service/internal/infraServiceConn/compression"
)

func TestCompressFilesSuccess(t *testing.T) {
	tempFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Creando archivo temporal: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString("contenido de prueba")
	if err != nil {
		t.Fatalf("Escribiendo en el archivo temporal: %v", err)
	}

	outputZip := "test.zip"
	defer os.Remove(outputZip)

	_, err = compression.CompressFiles(outputZip, []string{tempFile.Name()})
	if err != nil {
		t.Fatalf("Error al comprimir archivos: %v", err)
	}

	if _, err := os.Stat(outputZip); os.IsNotExist(err) {
		t.Fatalf("El archivo comprimido no existe: %v", err)
	}
}

func TestCompressFilesSpeed(t *testing.T) {
	tempFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Creando archivo temporal: %v", err)
	}
	defer os.Remove(tempFile.Name())

	outputZip := "test.zip"
	defer os.Remove(outputZip)

	startTime := time.Now()

	tempFiles := []string{}
	for i := 0; i < 10000; i++ {
		tempFiles = append(tempFiles, tempFile.Name())
	}

	_, err = compression.CompressFiles(
		outputZip,
		tempFiles,
	)
	if err != nil {
		t.Fatalf("Error al comprimir archivos: %v", err)
	}

	duration := time.Since(startTime)

	t.Logf("Comprimir el archivo tomó %v", duration)
}
