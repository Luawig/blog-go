package handler

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"blog-go/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateUser Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("CreateUser Error: %v", respData.Message)
	}
}

func TestGetUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateUser Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/user/1")
	if err != nil {
		t.Fatalf("GetUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetUser Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetUser Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetUser Error: %v", respData.Message)
	}

	userData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("GetUser Error: %v", "Data format error")
	}
	if userData["username"] != user.Username {
		t.Fatalf("GetUser Error: %v", "Data error")
	}
	if userData["email"] != user.Email {
		t.Fatalf("GetUser Error: %v", "Data error")
	}
}

func TestGetUserList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	for i := 0; i < 10; i++ {
		user := model.User{
			Username: "test" + strconv.Itoa(i),
			Password: "test",
			Email:    "test" + strconv.Itoa(i) + "@email.com",
		}

		userBytes, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("CreateUser Error: %v", err)
		}

		resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
		if err != nil {
			t.Fatalf("CreateUser Error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("CreateUser Error: %v", resp.Status)
		}
	}

	resp, err := http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/users?page_size=3&page_num=4")
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetUserList Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetUserList Error: %v", respData.Message)
	}

	userList, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetUserList Error: %v", "Data format error")
	}
	if len(userList) != 1 {
		t.Fatalf("GetUserList Error: %v", "Data error")
	}
}

func TestGetUserListByUsername(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	for i := 0; i < 10; i++ {
		user := model.User{
			Username: "test" + strconv.Itoa(i),
			Password: "test",
			Email:    "test" + strconv.Itoa(i) + "@email.com",
		}

		userBytes, err := json.Marshal(user)
		if err != nil {
			t.Fatalf("CreateUser Error: %v", err)
		}

		resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
		if err != nil {
			t.Fatalf("CreateUser Error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("CreateUser Error: %v", resp.Status)
		}
	}

	resp, err := http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/users/test")
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetUserList Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetUserList Error: %v", respData.Message)
	}

	userList, ok := respData.Data.([]interface{})
	if !ok {
		t.Fatalf("GetUserList Error: %v", "Data format error")
	}
	if len(userList) != 10 {
		t.Fatalf("GetUserList Error: %v", "Data error")
	}
}

func TestUpdateUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateUser Error: %v", resp.Status)
	}

	user.Email = "test2@email.com"
	userBytes, _ = json.Marshal(user)

	res, err := http.NewRequest(http.MethodPut, "http://localhost"+config.GetConfig().Server.Port+"/api/user/1", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("UpdateUser Error: %v", err)
	}
	resp, err = http.DefaultClient.Do(res)
	if err != nil {
		t.Fatalf("UpdateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateUser Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/user/1")
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetUserList Error: %v", resp.Status)
	}

	var respData utils.Response
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if respData.Status != utils.Success {
		t.Fatalf("GetUserList Error: %v", respData.Message)
	}

	userData, ok := respData.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("GetUserList Error: %v", "Data format error")
	}
	if userData["email"] != user.Email {
		t.Fatalf("GetUserList Error: %v", "Data error")
	}
}

func TestUpdateUserPassword(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}

	resp, err := http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("CreateUser Error: %v", resp.Status)
	}

	user = model.User{
		Password: "test2",
	}
	userBytes, _ = json.Marshal(user)

	res, err := http.NewRequest(http.MethodPut, "http://localhost"+config.GetConfig().Server.Port+"/api/user/1/password", bytes.NewReader(userBytes))
	if err != nil {
		t.Fatalf("UpdateUser Error: %v", err)
	}
	resp, err = http.DefaultClient.Do(res)
	if err != nil {
		t.Fatalf("UpdateUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateUser Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/user/1")
	if err != nil {
		t.Fatalf("GetUserList Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetUserList Error: %v", resp.Status)
	}
}

func TestDeleteUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()
	go routes.InitRouter()

	user := model.User{
		Username: "test",
		Password: "test",
		Email:    "test@email.com",
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("CreateUser Error: %v", err)
	}

	_, _ = http.Post("http://localhost"+config.GetConfig().Server.Port+"/api/user", "application/json", bytes.NewReader(userBytes))

	req, err := http.NewRequest(http.MethodDelete, "http://localhost"+config.GetConfig().Server.Port+"/api/user/1", nil)
	if err != nil {
		t.Fatalf("DeleteUser Error: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DeleteUser Error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("DeleteUser Error: %v", resp.Status)
	}

	resp, err = http.Get("http://localhost" + config.GetConfig().Server.Port + "/api/user/1")
	if err != nil {
		t.Fatalf("GetUser Error: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("GetUser Error: %v", resp.Status)
	}
}
