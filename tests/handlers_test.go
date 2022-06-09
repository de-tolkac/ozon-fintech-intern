package tests

import (
	"bytes"
	"encoding/json"
	"testing"
	. "github.com/de-tolkac/ozon-fintech-intern/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

var router *gin.Engine
var codeToUrl map[string]string

// Object to store JSON from POST /encode
type responseEncode struct {
	Code       int    `json:"code"`
	EncodedUrl string `json:"encodedUrl"`
	Err        string `json:"error"`
}

// Object to store JSON from GET /:short-url
type responseDecode struct {
	Code       int    `json:"code"`
	DecodedUrl string `json:"decodedUrl"`
	Err        string `json:"error"`
}

// Test object for generating short URL
type encodeTest struct {
	Req          []byte // In JSON format
	Url          string
	ResponseCode int
	Res          responseEncode
}

// Test object for finding long url by short
type decodeTest struct {
	Req          []byte // In JSON format
	Url          string
	ResponseCode int
	Res          responseDecode
}

func createRouter() {
	// If global config file wasn't previously generated
	if !ConfigInited {
		InitConfig()
	}

	router = gin.Default()
	router.POST("/encode", Encode(Cfg))
	router.GET("/:short-url", Decode(Cfg))
}

func TestEncode(t *testing.T) {
	// Call it only once in first test
	createRouter()

	var errorsTests = []encodeTest{
		{[]byte(`{"url" : "url.ru"}`), "url.ru", 200, responseEncode{1, "", "Invalid URL"}},
		{[]byte(`{"url" : "string"}`), "string", 200, responseEncode{1, "", "Invalid URL"}},
		{[]byte(`{"url" : "site.ru/param"}`), "site.ru/param", 200, responseEncode{1, "", "Invalid URL"}},
		{[]byte(`{"url" : ""}`), "", 400, responseEncode{2, "", "Invalid request body"}},
		{[]byte(`{"link" : "http://ya.ru"}`), "http://ya.ru", 400, responseEncode{2, "", "Invalid request body"}},
		{[]byte(`{}`), " ", 400, responseEncode{2, "", "Invalid request body"}},
	}

	for _, test := range errorsTests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(test.Req))
		router.ServeHTTP(w, req)

		assert.Equal(t, test.ResponseCode, w.Code)

		res := responseEncode{}
		err := json.Unmarshal([]byte(w.Body.String()), &res)
		assert.Equal(t, nil, err)
		assert.Equal(t, test.Res, res)
	}

	var successTests = []encodeTest{
		{[]byte(`{"url" : "http://ya.ru"}`), "http://ya.ru", 200, responseEncode{0, "", ""}},
		{[]byte(`{"url" : "httts://ozon.ru"}`), "httts://ozon.ru", 200, responseEncode{0, "", ""}},
		{[]byte(`{"url" : "httts://google.com/param"}`), "httts://google.com/param", 200, responseEncode{0, "", ""}},
	}

	codeToUrl = make(map[string]string)

	for _, test := range successTests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(test.Req))
		router.ServeHTTP(w, req)

		assert.Equal(t, test.ResponseCode, w.Code)

		res := responseEncode{}
		err := json.Unmarshal([]byte(w.Body.String()), &res)
		assert.Equal(t, nil, err)
		assert.Equal(t, test.Res.Err, res.Err)

		_, exists := codeToUrl[res.EncodedUrl]
		assert.Equal(t, false, exists)
		codeToUrl[res.EncodedUrl] = test.Url
	}

	/*
	// This test disabled because we don't need to truncate trailing slashes
	var sameUrlTests = []encodeTest{
		{[]byte(`{"url" : "http://ya.ru/"}`), "http://ya.ru", 200, responseEncode{0, "", ""}},
		{[]byte(`{"url" : "http://ya.ru////"}`), "http://ya.ru", 200, responseEncode{0, "", ""}},
		{[]byte(`{"url" : "http://ya.ru///"}`), "http://ya.ru", 200, responseEncode{0, "", ""}},
	}

	var prevUrl string
	for i := 0; i < len(sameUrlTests); i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(sameUrlTests[i].Req))
		router.ServeHTTP(w, req)

		assert.Equal(t, sameUrlTests[i].ResponseCode, w.Code)

		res := responseEncode{}
		err := json.Unmarshal([]byte(w.Body.String()), &res)
		assert.Equal(t, nil, err)
		assert.Equal(t, sameUrlTests[i].Res.Err, res.Err)

		if i != 0 {
			assert.Equal(t, prevUrl, res.EncodedUrl)
		} else {
			prevUrl = res.EncodedUrl
		}
	}
	*/
}

func TestDecode(t *testing.T) {
	for shortUrl, longUrl := range codeToUrl {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, shortUrl, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		res := responseDecode{}
		err := json.Unmarshal([]byte(w.Body.String()), &res)

		assert.Equal(t, nil, err)
		assert.Equal(t, "", res.Err)
		assert.Equal(t, longUrl, res.DecodedUrl)
	}
}
