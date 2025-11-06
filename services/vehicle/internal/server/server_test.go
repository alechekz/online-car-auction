package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alechekz/online-car-auction/services/vehicle/internal/server"
	"github.com/stretchr/testify/assert"
)

// TestNewServer_HandlerResponds checks that the server's handler responds to requests
func TestNewServer_HandlerResponds(t *testing.T) {
	cfg := server.NewConfig(":1111")
	srv := server.NewServer(cfg)

	req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
	rec := httptest.NewRecorder()

	srv.Handler().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
