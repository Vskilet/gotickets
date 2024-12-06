package handlers

import (
	"bytes"
	"encoding/json"
	"gotickets/db"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"gotickets/models"
)

func TestRegister(t *testing.T) {

	payload := models.UserRegister{
		FirstName: "Jean",
		LastName:  "BON",
		Email:     "jean@bon.br",
		Password:  "passw0rd",
	}
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	mok_db := db.NewDB()
	handler := NewHandler(mok_db)

	mockJsonPost(ctx, payload)
	handler.Register(ctx)

	payloadAnswer, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if ctx.Writer.Status() != http.StatusOK {
		t.Errorf("expected: %v, got: %v", http.StatusOK, ctx.Writer.Status())
	}

	// Check if the response body is correct.
	log.Println("body:", string(payloadAnswer))
}

func mockJsonPost(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
