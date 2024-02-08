package repository

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestCreateUser",
		Password: "TestPassword",
	}); code != utils.ErrorEmailEmpty {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Email:    "Test@email.com",
		Password: "TestPassword",
	}); code != utils.ErrorUsernameEmpty {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser",
		Email:    "Test@email.com",
	}); code != utils.ErrorPasswordEmpty {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser1",
		Email:    "Test1@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser2",
		Email:    "Test2@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser1",
		Email:    "Test@email.com",
		Password: "TestPassword",
	}); code != utils.ErrorUsernameUsed {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser",
		Email:    "Test1@email.com",
		Password: "TestPassword",
	}); code != utils.ErrorEmailUsed {
		t.Fatal("CreateUser failed")
	}
}

func TestGetUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestGetUser1",
		Email:    "Test1@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestGetUser2",
		Email:    "Test2@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}
	if user.Username != "TestGetUser1" {
		t.Fatal("GetUser failed")
	}
	if user.Email != "Test1@email.com" {
		t.Fatal("GetUser failed")
	}

	user, code = GetUser(2)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}
	if user.Username != "TestGetUser2" {
		t.Fatal("GetUser failed")
	}
	if user.Email != "Test2@email.com" {
		t.Fatal("GetUser failed")
	}

	if _, code = GetUser(3); code != utils.ErrorUserNotExist {
		t.Fatal("GetUser failed")
	}
}

func TestGetUserList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	for i := 0; i < 10; i++ {
		if code := CreateUser(&model.User{
			Username: "TestGetUserList" + strconv.Itoa(i),
			Email:    "Test" + strconv.Itoa(i) + "@email.com",
			Password: "TestPassword",
		}); code != utils.Success {
			t.Fatal("CreateUser failed")
		}
	}

	users, code := GetUserList(3, 2)
	if code != utils.Success {
		t.Fatal("GetUserList failed")
	}
	if len(users) != 3 {
		t.Fatal("GetUserList failed")
	}

	users, code = GetUserList(3, 4)
	if code != utils.Success {
		t.Fatal("GetUserList failed")
	}
	if len(users) != 1 {
		t.Fatal("GetUserList failed")
	}
}

func TestGetUserListByUsername(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	for i := 0; i < 10; i++ {
		if code := CreateUser(&model.User{
			Username: "TestUsername" + strconv.Itoa(i),
			Email:    "Test" + strconv.Itoa(i) + "@email.com",
			Password: "TestPassword",
		}); code != utils.Success {
			t.Fatal("CreateUser failed")
		}
	}

	for i := 0; i < 10; i++ {
		if code := CreateUser(&model.User{
			Username: "JustUsername" + strconv.Itoa(i),
			Email:    "Just" + strconv.Itoa(i) + "@email.com",
			Password: "TestPassword",
		}); code != utils.Success {
			t.Fatal("CreateUser failed")
		}
	}

	users, code := GetUserListByUsername("Test", 3, 2)
	if code != utils.Success {
		t.Fatal("GetUserList failed")
	}
	if len(users) != 3 {
		t.Fatal("GetUserList failed")
	}

	users, code = GetUserListByUsername("Test", 3, 4)
	if code != utils.Success {
		t.Fatal("GetUserList failed")
	}
	if len(users) != 1 {
		t.Fatal("GetUserList failed")
	}
}

func TestUpdateUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestCreateUser1",
		Email:    "Test1@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser2",
		Email:    "Test2@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := UpdateUser(1, &model.User{
		Username: "TestUpdateUser",
	}); code != utils.ErrorEmailEmpty {
		t.Fatal("UpdateUser failed")
	}

	if code := UpdateUser(1, &model.User{
		Email: "Test@email.com",
	}); code != utils.ErrorUsernameEmpty {
		t.Fatal("UpdateUser failed")
	}

	if code := UpdateUser(1, &model.User{
		Username: "TestCreateUser2",
		Email:    "Test@email.com",
		Password: "TestPassword",
	}); code != utils.ErrorUsernameUsed {
		t.Fatal("UpdateUser failed")
	}

	if code := UpdateUser(1, &model.User{
		Username: "TestUpdateUser",
		Email:    "Test2@email.com",
		Password: "TestPassword",
	}); code != utils.ErrorEmailUsed {
		t.Fatal("UpdateUser failed")
	}

	if code := UpdateUser(1, &model.User{
		Username: "TestUpdateUser",
		Email:    "Test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("UpdateUser failed")
	}
}

func TestUpdateUserPassword(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestUsername",
		Email:    "Test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := UpdateUserPassword(1, &model.User{
		Password: "TestPassword1",
	}); code != utils.Success {
		t.Fatal("UpdateUserPassword failed")
	}
}

func TestDeleteUser(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestCreateUser1",
		Email:    "Test1@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := CreateUser(&model.User{
		Username: "TestCreateUser2",
		Email:    "Test2@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	if code := DeleteUser(1); code != utils.Success {
		t.Fatal("DeleteUser failed")
	}

	if code := DeleteUser(1); code != utils.Success {
		t.Fatal("DeleteUser failed")
	}

	_, code := GetUser(1)
	if code != utils.ErrorUserNotExist {
		t.Fatal("DeleteUser failed")
	}
}

func TestGetUserWithPasswordByUsername(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "TestUsername",
		Email:    "TestEmail",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUserWithPasswordByUsername("TestUsername")
	if code != utils.Success {
		t.Fatal("GetUserPassword failed")
	}
	if user.Password != "TestPassword" {
		t.Fatal("GetUserPassword failed")
	}
}
