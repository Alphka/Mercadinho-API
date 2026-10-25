package routes

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func newHeadTestServer() *httptest.Server {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.HEAD("/explicit", func(context *gin.Context) {
		context.Status(http.StatusNoContent)
	})
	router.GET("/explicit", func(context *gin.Context) {
		context.String(http.StatusOK, "get")
	})

	// Like Swagger UI, rejects anything but GET
	router.GET("/get-only/*any", func(context *gin.Context) {
		if context.Request.Method != http.MethodGet {
			context.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		context.String(http.StatusOK, "body")
	})

	router.POST("/post-only", func(context *gin.Context) {
		context.Status(http.StatusCreated)
	})

	RegisterHeadRoutes(router)

	return httptest.NewServer(router)
}

// sendHead sends a raw HEAD request and returns the status line and any bytes
// received after the headers, which must be none.
func sendHead(t *testing.T, serverURL, path string) (status string, body []byte) {
	t.Helper()

	conn, err := net.Dial("tcp", strings.TrimPrefix(serverURL, "http://"))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("HEAD " + path + " HTTP/1.1\r\nHost: test\r\nConnection: close\r\n\r\n")); err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	status, err = reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
	}

	body, _ = io.ReadAll(reader)

	return strings.TrimSpace(status), body
}

func TestRegisterHeadRoutes(t *testing.T) {
	server := newHeadTestServer()
	defer server.Close()

	tests := []struct {
		name   string
		path   string
		status string
	}{
		{"mirrors GET route", "/get-only/index.html", "HTTP/1.1 200 OK"},
		{"keeps explicit HEAD route", "/explicit", "HTTP/1.1 204 No Content"},
		{"ignores non-GET route", "/post-only", "HTTP/1.1 404 Not Found"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status, body := sendHead(t, server.URL, test.path)

			if status != test.status {
				t.Errorf("status = %q, want %q", status, test.status)
			}

			// A leaked body would corrupt keep-alive connections
			if len(body) != 0 {
				t.Errorf("HEAD response has a %d byte body, want none", len(body))
			}
		})
	}
}
