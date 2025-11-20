package server_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alechekz/online-car-auction/services/pricing/internal/logger"
	"github.com/alechekz/online-car-auction/services/pricing/internal/server"
	"github.com/stretchr/testify/assert"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// TestNewServer_HandlerResponds checks that the demo server's handler responds to requests
func TestNewServer_HandlerResponds(t *testing.T) {
	cfg := server.NewConfig()
	srv, _ := server.NewServer(cfg)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Handler().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
