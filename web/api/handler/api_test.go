package apihandler_test

import (
	"github.com/devgek/webskeleton/web/echo"
	webenv "github.com/devgek/webskeleton/web/env"
	echofw "github.com/labstack/echo"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var authTokenAdminNoExpiresAt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwibmFtZSI6ImFkbWluIn0.BbGKax52n_5pqsutfKF62Gz6RdXHTJ9LPd9onWm8HuE"
var authString = "Bearer " + authTokenAdminNoExpiresAt
var echoForAPITests *echofw.Echo

/*
	Init initialize API tests
*/
func init() {
	echoForAPITests = echo.InitEchoApi(webenv.GetApiEnv(true))
}

func TestHandleAPIHealth(t *testing.T) {
	convey.Convey("Testing handler APIHealth", t, func() {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/api/health", nil)
		convey.So(err, convey.ShouldBeNil)

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		echoForAPITests.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		exp := `{"API":true,"Host":"","ProjectName":"webskeleton","VersionInfo":"V1.0","health":"ok"}`
		got := string(rr.Body.String())
		if strings.TrimRight(exp, "\n") != strings.TrimRight(got, "\n") {
			t.Errorf("handler returned unexpected body: got %v want %v",
				got, exp)
		}
	})
}

func TestHandleAPIEntityList(t *testing.T) {
	convey.Convey("Testing handler APIEntityList", t, func() {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/api/entitylistuser", nil)
		req.Header.Set("Authorization", authString)
		convey.So(err, convey.ShouldBeNil)

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		echoForAPITests.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		println(rr.Body.String())
	})
}

func TestHandleAPIOptionList(t *testing.T) {
	convey.Convey("Testing handler APIOptionList", t, func() {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/api/optionlistuser", nil)
		req.Header.Set("Authorization", authString)
		convey.So(err, convey.ShouldBeNil)

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		echoForAPITests.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		println(rr.Body.String())
	})
}
