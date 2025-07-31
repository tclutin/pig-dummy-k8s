package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      *API
		want       string
		statusCode int
	}{
		{
			name:       "check health",
			method:     http.MethodGet,
			input:      &API{},
			want:       `{"code":200,"message":"healthy"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with bad method",
			method:     http.MethodPost,
			input:      &API{},
			want:       `{"code":405,"message":"Method not allowed"}`,
			statusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/internal/healthz", nil)
			responseRecorder := httptest.NewRecorder()

			healthcheck(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}

}
