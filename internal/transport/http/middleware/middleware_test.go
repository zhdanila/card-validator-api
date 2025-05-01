package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoverMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(RecoverMiddleware())

	e.GET("/panic", func(c echo.Context) error {
		panic("something went wrong")
	})

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCORSMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(CORSMiddleware())

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "CORS OK")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set(echo.HeaderOrigin, "http://localhost:5173")
	req.Header.Set(echo.HeaderAccessControlRequestMethod, http.MethodGet)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "http://localhost:5173", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
	assert.Contains(t, rec.Header().Get(echo.HeaderAccessControlAllowMethods), http.MethodGet)
}
