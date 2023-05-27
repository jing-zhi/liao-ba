package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"community_id": 1,
		"title": "test",
		"content": "just a test"
    }`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// 判断响应的内容是不是按预期返回了需要登陆的的错误

	// 方法1： 判断相应内容是不是半酣指定的字符串
	//assert.Contains(t, w.Body.String(), "需要登录")
	// 方法2：将相应的内容反序列化到ResponseData 然后判断字段与预期是否一致
	res := new(ResponseDate)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body faild, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
