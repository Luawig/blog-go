package handler

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/routes"
	"blog-go/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

func TestCreateComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	comment := model.Comment{
		Content: "testComment",
		Article: &article,
		User:    &user,
	}
	commentBytes, err := json.Marshal(comment)
	if err != nil {
		t.Fatalf("CreateComment Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))
	if err != nil {
		t.Fatalf("CreateComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateComment Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("CreateComment Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("CreateComment Error: %v", respData.Message)
	}
}

func TestGetComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	comment := model.Comment{
		Content: "testComment",
		Article: &article,
		User:    &user,
	}
	commentBytes, _ := json.Marshal(comment)

	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/comment/1")
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetComment Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}

	commentData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
	if commentData["content"] != comment.Content {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
}

func TestGetCommentList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	for i := 0; i < 10; i++ {
		comment := model.Comment{
			Content: "testComment" + strconv.Itoa(i),
			Article: &article,
			User:    &user,
		}
		commentBytes, _ := json.Marshal(comment)

		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/comments?page_num=4&page_size=3")
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetComment Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}

	commentData, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
	if len(commentData) != 1 {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
}

func TestGetCommentListByArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	for i := 0; i < 10; i++ {
		comment := model.Comment{
			Content: "testComment" + strconv.Itoa(i),
			Article: &article,
			User:    &user,
		}
		commentBytes, _ := json.Marshal(comment)

		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/comments/article/1?page_num=4&page_size=3")
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetComment Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}

	commentData, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
	if len(commentData) != 1 {
		t.Fatalf("GetComment Error: %v", respData.Message)
	}
}

func TestUpdateComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	comment := model.Comment{
		Content: "testComment",
		Article: &article,
		User:    &user,
	}
	commentBytes, _ := json.Marshal(comment)

	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))

	resp, _ := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/login", "application/json", bytes.NewReader(userBytes))
	var respData utils.Response
	_ = json.NewDecoder(resp.Body).Decode(&respData)
	token, _ := respData.Data.(string)

	comment.Content = "testCommentUpdate"
	commentBytes, _ = json.Marshal(comment)
	req, err := http.NewRequest(http.MethodPut, "http://localhost"+config.GetServerConfig().Port+"/api/comment/1", bytes.NewReader(commentBytes))
	if err != nil {
		t.Fatalf("UpdateComment Error: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("UpdateComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateComment Error: %v", resp.Status)
	}

	resp, _ = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/comment/1")
	_ = json.NewDecoder(resp.Body).Decode(&respData)

	commentData, _ := respData.Data.(map[string]interface{})
	if commentData["content"] != comment.Content {
		t.Fatalf("UpdateComment Error: %v", respData.Message)
	}
}

func TestDeleteComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}
	userBytes, _ := json.Marshal(user)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	user.ID = 1

	article := model.Article{
		Title:   "test",
		Content: "test",
	}
	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	article.ID = 1

	comment := model.Comment{
		Content: "testComment",
		Article: &article,
		User:    &user,
	}
	commentBytes, _ := json.Marshal(comment)

	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/comment", "application/json", bytes.NewReader(commentBytes))

	resp, _ := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/login", "application/json", bytes.NewReader(userBytes))
	var respData utils.Response
	_ = json.NewDecoder(resp.Body).Decode(&respData)
	token, _ := respData.Data.(string)

	req, err := http.NewRequest(http.MethodDelete, "http://localhost"+config.GetServerConfig().Port+"/api/comment/1", nil)
	if err != nil {
		t.Fatalf("DeleteComment Error: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DeleteComment Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("DeleteComment Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/comment/1")
	if err != nil {
		t.Fatalf("GetComment Error: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("GetComment Error: %v", resp.Status)
	}
}
