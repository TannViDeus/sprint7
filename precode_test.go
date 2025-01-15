package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	const expectedCount = 4
	const count = 7

	url := fmt.Sprintf("/count?count=%d&city=%s", count, "moscow")
	req := httptest.NewRequest("GET", url, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	answer := responseRecorder.Body.String()
	answerLen := len(strings.Split(answer, ","))

	require.NotEmpty(t, answer)
	require.Equal(t, expectedCount, answerLen)
}

func TestMainHandlerWhenCityUnCorrect(t *testing.T) {
	const expectedCode = 400
	const expectedWrite = "wrong city value"
	const city = "omsk"

	url := fmt.Sprintf("/cafe?count=%d&city=%s", 4, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	answer := responseRecorder.Body.String()
	code := responseRecorder.Code

	assert.Equal(t, expectedCode, code)
	require.NotEmpty(t, answer)
	require.Equal(t, expectedWrite, answer)
}

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	const expectedCode = 200

	url := fmt.Sprintf("/cafe?count=%d&city=%s", 4, "moscow")
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	answer := responseRecorder.Body.String()
	code := responseRecorder.Code

	assert.Equal(t, expectedCode, code)
	assert.NotEmpty(t, answer)
}
