package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"fiber-example/internal/database"
	"fiber-example/internal/transport"
	"fiber-example/internal/models"
)

var dbName = "test.db"


type BookTestSuite struct {
	suite.Suite
	dbConn *gorm.DB
	app    *fiber.App
}


func (suite *BookTestSuite) SetupSuite() {

	// Удаление файла базы данных
	err := os.Remove(dbName)
	if err != nil {
		fmt.Println("Error removing database file:", err)
	} else {
		fmt.Println("Database file removed successfully")
	}

	suite.app = transport.Setup()
	database.InitDatabase(dbName)
	suite.dbConn = database.DBConn
}


func (suite *BookTestSuite) TearDownSuite() {
	
	// Закрытие соединения
	sqlDB, err := suite.dbConn.DB()
	if err != nil {
		panic("Failed to get DB from gorm")
	}
	err = sqlDB.Close()
	if err != nil {
		panic("Failed to close database connection")
	}
	fmt.Println("Database connection closed successfully")

	// Удаление файла базы данных
	err = os.Remove(dbName)
	if err != nil {
		fmt.Println("Error removing database file:", err)
	} else {
		fmt.Println("Database file removed successfully")
	}

	
}


func (suite *BookTestSuite) TestCreateBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": "12345", "title":"Test Book", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.app.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, res.StatusCode)

	var bookTest models.Book
	database.DBConn.Where("title = ?", "Test Book").First(&bookTest)
	fmt.Println(bookTest)
	assert.Equal(suite.T(), bookTest.Title, "Test Book")
}


func (suite *BookTestSuite) TestReadBook() {
	req := httptest.NewRequest(
		"GET",
		"/api/v1/book/12345",
		nil,
	)
	res, err := suite.app.Test(req, -1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var testbook models.Book
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &testbook)

	assert.Equal(suite.T(), "Test Book", testbook.Title)
}


func (suite *BookTestSuite) TestDeleteBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": "45678", "title":"Test Book", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.app.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, res.StatusCode)

	{
		req = httptest.NewRequest(
			"DELETE",
			"/api/v1/book/45678",
			nil,
		)
		res, err = suite.app.Test(req, -1)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

		var message Message
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &message)
		assert.Equal(suite.T(), "Book Successfully deleted", message.Msg)
	}

	{
		req = httptest.NewRequest(
			"DELETE",
			"/api/v1/book/45678",
			nil,
		)
		res, err = suite.app.Test(req, -1)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), http.StatusNotFound, res.StatusCode)
	
		var message Message
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &message)
		assert.Equal(suite.T(), "No Found The Book with ISIN", message.Msg)
	}


}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}

type Message struct {
	Msg string 
}