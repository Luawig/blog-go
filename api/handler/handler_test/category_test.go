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

func TestCreateCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}

	categoryBytes, err := json.Marshal(category)
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("CreateCategory Error: %v", respData.Message)
	}

	resp, err = http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("CreateCategory Error: %v", resp.Status)
	}
}

func TestGetCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}

	categoryBytes, err := json.Marshal(category)
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateCategory Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/category/1")
	if err != nil {
		t.Fatalf("GetCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetCategory Error: %v", respData.Message)
	}

	categoryData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("GetCategory Error: %v", "Data format error")
	}
	if categoryData["name"] != category.Name {
		t.Fatalf("GetCategory Error: %v", "Data error")
	}
}

func TestGetCategoryList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	for i := 0; i < 10; i++ {
		category := model.Category{
			Name: "test" + strconv.Itoa(i),
		}

		categoryBytes, err := json.Marshal(category)
		if err != nil {
			t.Fatalf("CreateCategory Error: %v", err)
		}

		resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
		if err != nil {
			t.Fatalf("CreateCategory Error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("CreateCategory Error: %v", resp.Status)
		}
	}

	resp, err := http.Get("http://localhost" + config.GetServerConfig().Port + "/api/categories")
	if err != nil {
		t.Fatalf("GetCategoryList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetCategoryList Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetCategoryList Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetCategoryList Error: %v", respData.Message)
	}

	categoryList, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetCategoryList Error: %v", "Data format error")
	}
	if len(categoryList) != 10 {
		t.Fatalf("GetCategoryList Error: %v", "Data error")
	}
}

func TestUpdateCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}

	categoryBytes, err := json.Marshal(category)
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateCategory Error: %v", resp.Status)
	}

	category.Name = "test1"
	categoryBytes, _ = json.Marshal(category)

	req, err := http.NewRequest(http.MethodPut, "http://localhost"+config.GetServerConfig().Port+"/api/category/1", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("UpdateCategory Error: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("UpdateCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateCategory Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/category/1")
	if err != nil {
		t.Fatalf("GetCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("UpdateCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("UpdateCategory Error: %v", respData.Message)
	}

	categoryData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("UpdateCategory Error: %v", "Data format error")
	}
	if categoryData["name"] != category.Name {
		t.Fatalf("UpdateCategory Error: %v", "Data error")
	}
}

func TestDeleteCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	category := model.Category{
		Name: "test",
	}

	categoryBytes, err := json.Marshal(category)
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetServerConfig().Port+"/api/category", "application/json", bytes.NewReader(categoryBytes))
	if err != nil {
		t.Fatalf("CreateCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateCategory Error: %v", resp.Status)
	}

	req, err := http.NewRequest(http.MethodDelete, "http://localhost"+config.GetServerConfig().Port+"/api/category/1", nil)
	if err != nil {
		t.Fatalf("DeleteCategory Error: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DeleteCategory Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("DeleteCategory Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("DeleteCategory Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("DeleteCategory Error: %v", respData.Message)
	}

	resp, err = http.Get("http://localhost" + config.GetServerConfig().Port + "/api/category/1")
	if err != nil {
		t.Fatalf("GetCategory Error: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("DeleteCategory Error: %v", resp.Status)
	}
}
