package db

import (
	"blog-go/config"
	"blog-go/internal/model"
	"fmt"
	"testing"
	"time"
)

func TestInitDB(t *testing.T) {
	config.InitConfig()
	InitDB()

	// Add two categories
	category1 := model.Category{Name: "CategoryTest1"}
	category2 := model.Category{Name: "CategoryTest2"}

	db.Create(&category1)
	db.Create(&category2)

	// Query categories
	var categories []model.Category
	db.Preload("Articles").Find(&categories)
	if len(categories) != 2 {
		t.Errorf("Query categories failed")
	}
	fmt.Println(categories)

	// Add two articles
	article1 := model.Article{
		Title:        "TitleTest1",
		Content:      "ContentTest1",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		CommentCount: 0,
		ReadCount:    0,
		Categories:   []*model.Category{&category1},
	}
	time.Sleep(1 * time.Second)
	article2 := model.Article{
		Title:        "TitleTest2",
		Content:      "ContentTest2",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		CommentCount: 0,
		ReadCount:    0,
		Categories:   []*model.Category{&category1, &category2},
	}

	db.Create(&article1)
	db.Create(&article2)

	// Query articles
	var articles []model.Article
	db.Preload("Categories").Preload("Comments").Find(&articles)
	if len(articles) != 2 {
		t.Errorf("Query articles failed")
	}
	fmt.Println(articles)

	// Add a user
	user1 := model.User{
		Username:      "UsernameTest1",
		Password:      "PasswordTest1",
		Email:         "Email1@Test.com",
		CreateTime:    time.Now(),
		LastLoginTime: time.Now(),
	}
	db.Create(&user1)

	// Query users
	var users []model.User
	db.Preload("Comments").Find(&users)
	if len(users) != 1 {
		t.Errorf("Query users failed")
	}
	fmt.Println(users)

	// Add two comments
	comment1 := model.Comment{
		Content: "ContentTest1",
		Time:    time.Now(),
		Article: article1,
		User:    user1,
	}
	time.Sleep(1 * time.Second)
	comment2 := model.Comment{
		Content: "ContentTest2",
		Time:    time.Now(),
		Article: article2,
		User:    user1,
	}
	db.Create(&comment1)
	db.Create(&comment2)

	// Query comments
	var comments []model.Comment
	db.Preload("Article").Preload("User").Find(&comments)
	if len(comments) != 2 {
		t.Errorf("Query comments failed")
	}
	fmt.Println(comments)
}
