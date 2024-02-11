package handler

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"blog-go/routes"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "TestUsername",
		Password: "TestPassword",
		Email:    "Test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))

	resp, _ := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/login", "application/json", bytes.NewReader(userBytes))
	var respData utils.Response
	_ = json.NewDecoder(resp.Body).Decode(&respData)
	token := respData.Data.(string)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "static/favicon.ico")
	file, err := os.Open("static/favicon.ico")
	_, _ = io.Copy(part, file)
	_ = writer.Close()

	req, err := http.NewRequest("POST", "http://localhost"+config.GetServerConfig().Port+"/api/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code 200, but got %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatal(err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("expected code 200, but %s", respData.Message)
	}

	url, ok := respData.Data.(string)
	if !ok {
		t.Fatalf("expected string, but %T", respData.Data)
	}
	_ = url
}
