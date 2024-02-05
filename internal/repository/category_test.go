package repository

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/internal/model"
	"blog-go/pkg/utils"
	"fmt"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{}); code != utils.ErrorCategoryNameEmpty {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestCreateCategory1",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestCreateCategory2",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestCreateCategory1",
	}); code != utils.ErrorCategoryNameUsed {
		t.Fatal("CreateCategory failed")
	}
}

func TestGetCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{
		Name: "TestGetCategory1",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestGetCategory2",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}
	if category.Name != "TestGetCategory1" {
		t.Fatal("GetCategory failed")
	}

	category, code = GetCategory(2)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}
	if category.Name != "TestGetCategory2" {
		t.Fatal("GetCategory failed")
	}

	if _, code = GetCategory(3); code != utils.ErrorCategoryNotExist {
		t.Fatal("GetCategory failed")
	}
}

func TestGetCategoryList(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{
		Name: "TestGetCategoryList1",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestGetCategoryList2",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	categories, code := GetCategoryList()
	if code != utils.Success {
		t.Fatal("GetCategoryList failed")
	}
	if len(categories) != 2 {
		fmt.Printf("%+v\n", categories)
		t.Fatal("GetCategoryList failed")
	}
	fmt.Printf("%+v\n", categories)
}

func TestUpdateCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{
		Name: "TestUpdateCategory1",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestUpdateCategory2",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := UpdateCategory(1, &model.Category{
		Name: "TestUpdateCategory3",
	}); code != utils.Success {
		t.Fatal("UpdateCategory failed")
	}

	category, code := GetCategory(1)
	if code != utils.Success {
		t.Fatal("GetCategory failed")
	}
	if category.Name != "TestUpdateCategory3" {
		t.Fatal("UpdateCategory failed")
	}

	if code := UpdateCategory(1, &model.Category{
		Name: "TestUpdateCategory2",
	}); code != utils.ErrorCategoryNameUsed {
		t.Fatal("UpdateCategory failed")
	}
}

func TestDeleteCategory(t *testing.T) {
	config.InitConfig()
	db.InitTestDB()

	if code := CreateCategory(&model.Category{
		Name: "TestDeleteCategory1",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := CreateCategory(&model.Category{
		Name: "TestDeleteCategory2",
	}); code != utils.Success {
		t.Fatal("CreateCategory failed")
	}

	if code := DeleteCategory(1); code != utils.Success {
		t.Fatal("DeleteCategory failed")
	}

	if code := DeleteCategory(1); code != utils.ErrorCategoryNotExist {
		t.Fatal("DeleteCategory failed")
	}

	category, code := GetCategory(1)
	if code != utils.ErrorCategoryNotExist {
		t.Fatal("DeleteCategory failed")
	}
	if category != nil {
		t.Fatal("DeleteCategory failed")
	}
}
