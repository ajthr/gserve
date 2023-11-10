package server

import (
    "net/http"
    "net/http/httptest"
    "testing"

	"github.com/stretchr/testify/assert"
)

func TestServerHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
    indexHTMLTemplateHandler(w, req)
    res := w.Result()

	assert.Equal(t, 200, res.StatusCode)

}
