package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yhlin66/go-side/controller"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomepageHandler(t *testing.T) {
	//預計拿到的值
	mockResponse := `{"message":"Here is HomePage"}`
	r := SetUpRouter()
	r.GET("/", controller.HomepageHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	//紀錄 request 
	w := httptest.NewRecorder()
	// handler 啟動 serverHTTP
	r.ServeHTTP(w, req)

	//回傳值
	responseData, _ := ioutil.ReadAll(w.Body)
	//測試 responseData 是否與預計的值 相等
	assert.Equal(t, mockResponse, string(responseData))
	// 測試 statusCode 
	assert.Equal(t, http.StatusOK, w.Code)
}