package handler_test

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

func TestCreateArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	article := model.Article{
		Title:   "test",
		Content: "test",
	}

	articleBytes, err := json.Marshal(article)
	if err != nil {
		t.Fatalf("CreateArticle Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	if err != nil {
		t.Fatalf("CreateArticle Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateArticle Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("CreateArticle Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("CreateArticle Error: %v", respData.Message)
	}
}

func TestGetArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	article := model.Article{
		Title:   "test",
		Content: "test",
	}

	articleBytes, err := json.Marshal(article)
	if err != nil {
		t.Fatalf("CreateArticle Error: %v", err)
	}

	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/article/1")
	if err != nil {
		t.Fatalf("GetArticle Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetArticle Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetArticle Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetArticle Error: %v", respData.Message)
	}

	articleData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("GetArticle Error: %v", respData.Message)
	}
	if articleData["title"] != article.Title {
		t.Fatalf("GetArticle Error: %v", "title not equal")
	}
	if articleData["content"] != article.Content {
		t.Fatalf("GetArticle Error: %v", "content not equal")
	}
}

func TestGetArticleList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	for i := 0; i < 10; i++ {
		article := model.Article{
			Title:   "test" + strconv.Itoa(i),
			Content: "test" + strconv.Itoa(i),
		}
		articleBytes, _ := json.Marshal(article)
		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/articles?page_num=4&page_size=3")
	if err != nil {
		t.Fatalf("GetArticleList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetArticleList Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetArticleList Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetArticleList Error: %v", respData.Message)
	}

	articleData, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetArticleList Error: %v", respData.Message)
	}
	if len(articleData) != 1 {
		t.Fatalf("GetArticleList Error: %v", "article list is empty")
	}
}

func TestGetArticleListByCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}
	categoryBytes, _ := json.Marshal(category)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	category.ID = 1

	for i := 0; i < 10; i++ {
		article := model.Article{
			Title:      "test" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{&category},
		}
		articleBytes, _ := json.Marshal(article)
		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/articles/category/1")
	if err != nil {
		t.Fatalf("GetArticleListByCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetArticleListByCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetArticleListByCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetArticleListByCategory Error: %v", respData.Message)
	}

	articleData, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetArticleListByCategory Error: %v", respData.Message)
	}
	if len(articleData) != 10 {
		t.Fatalf("GetArticleListByCategory Error: %v", "article list is empty")
	}
}

func TestGetArticleListByTitle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}
	categoryBytes, _ := json.Marshal(category)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	category.ID = 1

	for i := 0; i < 10; i++ {
		article := model.Article{
			Title:      "test" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{&category},
		}
		articleBytes, _ := json.Marshal(article)
		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	}

	for i := 0; i < 10; i++ {
		article := model.Article{
			Title:      "title" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{&category},
		}
		articleBytes, _ := json.Marshal(article)
		_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/articles/test?page_num=2&page_size=3")
	if err != nil {
		t.Fatalf("GetArticleListByCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetArticleListByCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetArticleListByCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetArticleListByCategory Error: %v", respData.Message)
	}

	articleData, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetArticleListByCategory Error: %v", respData.Message)
	}
	if len(articleData) != 3 {
		t.Fatalf("GetArticleListByCategory Error: %v", "article list is wrong")
	}
}

func TestUpdateArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	article := model.Article{
		Title:   "test",
		Content: "test",
	}

	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))

	article.Title = "test1"
	articleBytes, _ = json.Marshal(article)
	req, _ := http.NewRequest("PUT", "http://localhost"+config.GetServerConfig().Port+"/api/article/1", bytes.NewReader(articleBytes))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("UpdateArticle Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateArticle Error: %v", resp.Status)
	}

	resp, _ = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/article/1")

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("UpdateArticle Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("UpdateArticle Error: %v", respData.Message)
	}

	articleData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("UpdateArticle Error: %v", respData.Message)
	}
	if articleData["title"] != article.Title {
		t.Fatalf("UpdateArticle Error: %v", "title not equal")
	}
}

func TestDeleteArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	article := model.Article{
		Title:   "test",
		Content: "test",
	}

	articleBytes, _ := json.Marshal(article)
	_, _ = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/article", "application/json", bytes.NewReader(articleBytes))

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost"+config.GetServerConfig().Port+"/api/article/1", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DeleteArticle Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("DeleteArticle Error: %v", resp.Status)
	}

	resp, _ = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/article/1")

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("DeleteArticle Error: %v", err)
	}
	if respData.Status == utils.Success {
		t.Fatalf("DeleteArticle Error: %v", respData.Message)
	}
}
