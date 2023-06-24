package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )

// const UploadDirectory = "./data"

// func TestHomePage(t *testing.T) {
// 	_, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestListFilesIntegrity(t *testing.T) {
// 	_, err := http.NewRequest("GET", "/files", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestUploadFile(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(UploadFile))
// 	defer ts.Close()

// 	file, err := ioutil.TempFile("", "test-file")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.Remove(file.Name())

// 	body := strings.NewReader("test content")
// 	req, err := http.NewRequest("POST", ts.URL, body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "multipart/form-data")
// 	req.Header.Set("Content-Disposition", "form-data; name=file; filename=test-file")
// 	req.ContentLength = int64(body.Len())

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", res.StatusCode)
// 	}

// 	_, err = os.Stat(filepath.Join(BaseDir, "test-file"))
// 	if err != nil {
// 		t.Errorf("failed to create the file: %v", err)
// 	}
// }

// func TestDownloadFile(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(DownloadFile))
// 	defer ts.Close()

// 	filePath := filepath.Join(BaseDir, "test-file")
// 	err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer os.Remove(filePath)

// 	req, err := http.NewRequest("GET", ts.URL+"/test-file", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", res.StatusCode)
// 	}

// 	content, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expectedContent := "test content"
// 	if string(content) != expectedContent {
// 		t.Errorf("expected content %q; got %q", expectedContent, string(content))
// 	}
// }

// func TestDeleteFile(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(DeleteFile))
// 	defer ts.Close()

// 	filePath := filepath.Join(BaseDir, "test-file")
// 	err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, err := http.NewRequest("DELETE", ts.URL+"/test-file", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", res.StatusCode)
// 	}

// 	_, err = os.Stat(filePath)
// 	if !os.IsNotExist(err) {
// 		t.Errorf("expected file to be deleted; got %v", err)
// 	}
// }

// func TestListFiles(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(ListFiles))
// 	defer ts.Close()

// 	filePaths := []string{
// 		filepath.Join(BaseDir, "file1.txt"),
// 		filepath.Join(BaseDir, "file2.txt"),
// 	}
// 	for _, filePath := range filePaths {
// 		err := ioutil.WriteFile(filePath, []byte("test content"), 0644)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer os.Remove(filePath)
// 	}

// 	req, err := http.NewRequest("GET", ts.URL, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status OK; got %v", res.StatusCode)
// 	}

// 	var fileInfos []FileInfo
// 	err = json.NewDecoder(res.Body).Decode(&fileInfos)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, filePath := range filePaths {
// 		found := false
// 		for _, fileInfo := range fileInfos {
// 			if fileInfo.Name == filepath.Base(filePath) {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			t.Errorf("expected file %q to be listed; not found", filePath)
// 		}
// 	}
// }
