package repository

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"strconv"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

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
}

func TestGetArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

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

	if _, code := GetArticle(1); code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if _, code := GetArticle(2); code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	if _, code := GetArticle(3); code != utils.ErrorArticleNotExist {
		t.Fatal("GetArticle failed")
	}
}

func TestGetArticleList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	db.DB.Create(&model.Category{
		Name: "test",
	})
	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}

	for i := 0; i < 10; i++ {
		if code := CreateArticle(&model.Article{
			Title:      "test" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{category},
		}); code != utils.Success {
			t.Fatal("CreateArticle failed")
		}
	}

	articles, code := GetArticleList(3, 2)
	if code != utils.Success {
		t.Fatal("GetArticleList failed")
	}
	if len(articles) != 3 {
		t.Fatal("GetArticleList failed")
	}

	articles, code = GetArticleList(3, 4)
	if code != utils.Success {
		t.Fatal("GetArticleList failed")
	}
	if len(articles) != 1 {
		t.Fatal("GetArticleList failed")
	}
}

func TestGetArticleListByCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	db.DB.Create(&model.Category{
		Name: "test",
	})
	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}

	for i := 0; i < 10; i++ {
		if code := CreateArticle(&model.Article{
			Title:      "test" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{category},
		}); code != utils.Success {
			t.Fatal("CreateArticle failed")
		}
	}

	articles, code := GetArticleListByCategory(1, 3, 2)
	if code != utils.Success {
		t.Fatal("GetArticleListByCategory failed")
	}
	if len(articles) != 3 {
		t.Fatal("GetArticleListByCategory failed")
	}

	articles, code = GetArticleListByCategory(1, 3, 4)
	if code != utils.Success {
		t.Fatal("GetArticleListByCategory failed")
	}
	if len(articles) != 1 {
		t.Fatal("GetArticleListByCategory failed")
	}

	articles, code = GetArticleListByCategory(2, 3, 2)
	if code != utils.Success {
		t.Fatal("GetArticleListByCategory failed")
	}
	if len(articles) != 0 {
		t.Fatal("GetArticleListByCategory failed")
	}
}

func TestGetArticleListByTitle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	db.DB.Create(&model.Category{
		Name: "test",
	})
	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}

	for i := 0; i < 10; i++ {
		if code := CreateArticle(&model.Article{
			Title:      "test" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{category},
		}); code != utils.Success {
			t.Fatal("CreateArticle failed")
		}
	}
	for i := 0; i < 10; i++ {
		if code := CreateArticle(&model.Article{
			Title:      "title" + strconv.Itoa(i),
			Content:    "test" + strconv.Itoa(i),
			Categories: []*model.Category{category},
		}); code != utils.Success {
			t.Fatal("CreateArticle failed")
		}
	}

	articles, code := GetArticleListByTitle("test", 3, 2)
	if code != utils.Success {
		t.Fatal("GetArticleListByTitle failed")
	}
	if len(articles) != 3 {
		t.Fatal("GetArticleListByTitle failed")
	}

	articles, code = GetArticleListByTitle("test", 3, 4)
	if code != utils.Success {
		t.Fatal("GetArticleListByTitle failed")
	}
	if len(articles) != 1 {
		t.Fatal("GetArticleListByTitle failed")
	}
}

func TestUpdateArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

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

	articles, code := GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}

	articles.Title = "test3"
	if code := UpdateArticle(articles); code != utils.Success {
		t.Fatal("UpdateArticle failed")
	}

	articles, code = GetArticle(1)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}
	if articles.Title != "test3" {
		t.Fatal("UpdateArticle failed")
	}
}

func TestDeleteArticle(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{
		Name: "test",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}

	if code := CreateArticle(&model.Article{
		Title:      "test1",
		Content:    "test1",
		Categories: []*model.Category{category},
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	if code := CreateArticle(&model.Article{
		Title:      "test2",
		Content:    "test2",
		Categories: []*model.Category{category},
	}); code != utils.Success {
		t.Fatal("CreateArticle failed")
	}

	if code := DeleteArticle(1); code != utils.Success {
		t.Fatal("DeleteArticle failed")
	}

	if _, code := GetArticle(1); code != utils.ErrorArticleNotExist {
		t.Fatal("DeleteArticle failed")
	}

	if code := DeleteCategory(1); code != utils.Success {
		t.Fatal("DeleteCategory failed")
	}

	article, code := GetArticle(2)
	if code != utils.Success {
		t.Fatal("GetArticle failed")
	}
	if len(article.Categories) != 0 {
		t.Fatal("DeleteCategory failed")
	}
}
