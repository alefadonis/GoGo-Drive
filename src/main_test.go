package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const UploadDirectory = "./data"

func TestHomePage(t *testing.T) {
	_, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListFilesIntegrity(t *testing.T) {
	_, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadFile(t *testing.T) {
	// Preparar um servidor de teste
	ts := httptest.NewServer(http.HandlerFunc(UploadFile))
	defer ts.Close()

	// Criar um arquivo de teste
	file, err := ioutil.TempFile("", "test-file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Criar uma solicitação HTTP POST para enviar o arquivo
	body := strings.NewReader("test content")
	req, err := http.NewRequest("POST", ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Disposition", "form-data; name=file; filename=test-file")
	req.ContentLength = int64(body.Len())

	// Enviar a solicitação para o servidor de teste
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Verificar o código de status da resposta
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	// Verificar se o arquivo foi criado com sucesso
	_, err = os.Stat(filepath.Join(BaseDir, "test-file"))
	if err != nil {
		t.Errorf("failed to create the file: %v", err)
	}
}

func TestDownloadFile(t *testing.T) {
	// Preparar um servidor de teste
	ts := httptest.NewServer(http.HandlerFunc(DownloadFile))
	defer ts.Close()

	// Criar um arquivo de teste
	filePath := filepath.Join(BaseDir, "test-file")
	err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	// Criar uma solicitação HTTP GET para baixar o arquivo
	req, err := http.NewRequest("GET", ts.URL+"/test-file", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enviar a solicitação para o servidor de teste
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Verificar o código de status da resposta
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	// Ler o conteúdo do arquivo baixado
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Verificar se o conteúdo do arquivo está correto
	expectedContent := "test content"
	if string(content) != expectedContent {
		t.Errorf("expected content %q; got %q", expectedContent, string(content))
	}
}

func TestDeleteFile(t *testing.T) {
	// Preparar um servidor de teste
	ts := httptest.NewServer(http.HandlerFunc(DeleteFile))
	defer ts.Close()

	// Criar um arquivo de teste
	filePath := filepath.Join(BaseDir, "test-file")
	err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Criar uma solicitação HTTP DELETE para excluir o arquivo
	req, err := http.NewRequest("DELETE", ts.URL+"/test-file", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enviar a solicitação para o servidor de teste
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Verificar o código de status da resposta
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	// Verificar se o arquivo foi excluído com sucesso
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("expected file to be deleted; got %v", err)
	}
}

func TestListFiles(t *testing.T) {
	// Preparar um servidor de teste
	ts := httptest.NewServer(http.HandlerFunc(ListFiles))
	defer ts.Close()

	// Criar alguns arquivos de teste
	filePaths := []string{
		filepath.Join(BaseDir, "file1.txt"),
		filepath.Join(BaseDir, "file2.txt"),
	}
	for _, filePath := range filePaths {
		err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(filePath)
	}

	// Criar uma solicitação HTTP GET para listar os arquivos
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enviar a solicitação para o servidor de teste
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Verificar o código de status da resposta
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	// Ler a resposta JSON
	var fileInfos []FileInfo
	err = json.NewDecoder(res.Body).Decode(&fileInfos)
	if err != nil {
		t.Fatal(err)
	}

	// Verificar se os arquivos estão presentes na resposta
	for _, filePath := range filePaths {
		found := false
		for _, fileInfo := range fileInfos {
			if fileInfo.Name == filepath.Base(filePath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected file %q to be listed; not found", filePath)
		}
	}
}
