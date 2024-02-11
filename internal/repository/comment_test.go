package repository

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/utils"
	"strconv"
	"testing"
)

func TestCreateComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test",
		Content: "test",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test",
		User:    user,
		Article: article,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}
}

func TestGetComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test",
		Content: "test",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test",
		User:    user,
		Article: article,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	comment, code := GetComment(1)
	if code != utils.Success {
		t.Fatal("GetComment failed")
	}
	if comment == nil {
		t.Fatal("GetComment failed")
	}
}

func TestGetCommentList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test",
		Content: "test",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	for i := 0; i < 10; i++ {
		if code := CreateComment(&model.Comment{
			Content: "test" + strconv.Itoa(i),
			User:    user,
			Article: article,
		}); code != utils.Success {
			t.Fatal("CreateComment failed")
		}
	}

	comments, code := GetCommentList(3, 2)
	if code != utils.Success {
		t.Fatal("GetCommentList failed")
	}
	if len(comments) != 3 {
		t.Fatal("GetCommentList failed")
	}

	comments, code = GetCommentList(3, 4)
	if code != utils.Success {
		t.Fatal("GetCommentList failed")
	}
	if len(comments) != 1 {
		t.Fatal("GetCommentList failed")
	}
}

func TestGetCommentListByArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test1",
		Content: "test1",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test2",
		Content: "test2",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article1, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	article2, code := GetArticle(2)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	for i := 0; i < 10; i++ {
		if code := CreateComment(&model.Comment{
			Content: "test1" + strconv.Itoa(i),
			User:    user,
			Article: article1,
		}); code != utils.Success {
			t.Fatal("CreateComment failed")
		}
		if code := CreateComment(&model.Comment{
			Content: "test2" + strconv.Itoa(i),
			User:    user,
			Article: article2,
		}); code != utils.Success {
			t.Fatal("CreateComment failed")
		}
	}

	comments, code := GetCommentListByArticle(1, 3, 2)
	if code != utils.Success {
		t.Fatal("GetCommentListByArticle failed")
	}
	if len(comments) != 3 {
		t.Fatal("GetCommentListByArticle failed")
	}

	comments, code = GetCommentListByArticle(1, 3, 4)
	if code != utils.Success {
		t.Fatal("GetCommentListByArticle failed")
	}
	if len(comments) != 1 {
		t.Fatal("GetCommentListByArticle failed")
	}
}

func TestGetCommentUserID(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test",
		Content: "test",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test",
		User:    user,
		Article: article,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	userId, code := GetCommentUserID(1)
	if code != utils.Success {
		t.Fatal("GetCommentUserID failed")
	}
	if userId != 1 {
		t.Fatal("GetCommentUserID failed")
	}
}

func TestUpdateComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test",
		Content: "test",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test",
		User:    user,
		Article: article,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	comment, code := GetComment(1)
	if code != utils.Success {
		t.Fatal("GetComment failed")
	}
	if comment == nil {
		t.Fatal("GetComment failed")
	}

	comment.Content = "test2"
	if code := UpdateComment(1, comment); code != utils.Success {
		t.Fatal("UpdateComment failed")
	}

	comment, code = GetComment(1)
	if code != utils.Success {
		t.Fatal("GetComment failed")
	}
	if comment.Content != "test2" {
		t.Fatal("UpdateComment failed")
	}
}

func TestDeleteComment(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateUser(&model.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "TestPassword",
	}); code != utils.Success {
		t.Fatal("CreateUser failed")
	}

	user, code := GetUser(1)
	if code != utils.Success {
		t.Fatal("GetUser failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test1",
		Content: "test1",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	if code := CreateArticle(&model.Article{
		Title:   "test2",
		Content: "test2",
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	article1, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	article2, code := GetArticle(2)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test1",
		User:    user,
		Article: article1,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test2",
		User:    user,
		Article: article1,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	if code := CreateComment(&model.Comment{
		Content: "test3",
		User:    user,
		Article: article2,
	}); code != utils.Success {
		t.Fatal("CreateComment failed")
	}

	if code := DeleteComment(1); code != utils.Success {
		t.Fatal("DeleteComment failed")
	}

	if _, code := GetComment(1); code == utils.Success {
		t.Fatal("DeleteComment failed")
	}

	if code := DeleteArticle(1); code != utils.Success {
		t.Fatal("DeleteUser failed")
	}

	if _, code := GetComment(2); code == utils.Success {
		t.Fatal("DeleteComment failed")
	}

	if code := DeleteUser(1); code != utils.Success {
		t.Fatal("DeleteUser failed")
	}

	if _, code := GetComment(3); code == utils.Success {
		t.Fatal("DeleteComment failed")
	}
}
