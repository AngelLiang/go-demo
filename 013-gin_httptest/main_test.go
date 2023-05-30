package main

import (
	// "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
    router := gin.Default()
	InitRouter(router)

    client := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/v1/helloworld", nil)
    router.ServeHTTP(client, req)

	// fmt.Println(client)
    assert.Equal(t, 200, client.Code)
    assert.Equal(t, "\"helloworld\"", client.Body.String())
}
