package adapters

import (
	"bff/domain"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestApigeeAdapter_GetEnv(t *testing.T) {
	type fields struct {
		environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"it works",
			fields{
				environment: "TEST",
			},
			"TEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApigeeAdapter{
				environment: tt.fields.environment,
			}
			if got := a.GetEnv(); got != tt.want {
				t.Errorf("ApigeeAdapter.GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApigeeAdapter_DoRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		c.String(200, "Hello")
	})
	router.POST("/post", func(c *gin.Context) {
		c.String(200, "Hello")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	type fields struct {
		endpoint     string
		clientID     string
		clientSecret string
		environment  string
		authPath     string
		httpClient   *http.Client
	}
	type args struct {
		method string
		url    string
		header *http.Header
		body   io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		{
			"test on invalid method returns error",
			fields{
				endpoint:   "test/test",
				httpClient: &http.Client{},
			},
			args{method: "/INVALID"},
			0,
			nil,
			true,
		},
		{
			"test on invalid endpoint format returns error",
			fields{
				endpoint:   "",
				httpClient: &http.Client{},
			},
			args{
				header: &http.Header{},
				url:    "",
			},
			0,
			nil,
			true,
		},
		{
			"Test GET runs OK",
			fields{
				endpoint:   ts.URL,
				httpClient: &http.Client{},
			},
			args{
				header: &http.Header{},
				url:    "/get",
			},
			200,
			[]byte("Hello"),
			false,
		},
		{
			"Test POST runs OK",
			fields{
				endpoint:   ts.URL,
				httpClient: &http.Client{},
			},
			args{
				method: "POST",
				header: &http.Header{},
				url:    "/post",
			},
			200,
			[]byte("Hello"),
			false,
		},
		{
			"Test non-existent url gets 404",
			fields{
				endpoint:   ts.URL,
				httpClient: &http.Client{},
			},
			args{
				header: &http.Header{},
				url:    "/asdasdas",
			},
			404,
			[]byte("404 page not found"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApigeeAdapter{
				endpoint:     tt.fields.endpoint,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				environment:  tt.fields.environment,
				authPath:     tt.fields.authPath,
				httpClient:   tt.fields.httpClient,
			}
			got, got1, err := a.DoRequest(tt.args.method, tt.args.url, tt.args.header, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApigeeAdapter.DoRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ApigeeAdapter.DoRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ApigeeAdapter.DoRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewApigeeAdapter(t *testing.T) {
	type args struct {
		endpoint     string
		clientID     string
		clientSecret string
		environment  string
		authPath     string
	}
	tests := []struct {
		name    string
		args    args
		want    *ApigeeAdapter
		wantErr bool
	}{
		{
			"test runs correctly",
			args{
				endpoint:     "endpoint",
				clientID:     "clientID",
				clientSecret: "clientSecret",
				environment:  "environment",
				authPath:     "authPath",
			},
			&ApigeeAdapter{
				endpoint:     "endpoint",
				clientID:     "clientID",
				clientSecret: "clientSecret",
				environment:  "environment",
				authPath:     "authPath",
				httpClient:   &http.Client{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewApigeeAdapter(tt.args.endpoint, tt.args.clientID, tt.args.clientSecret, tt.args.environment, tt.args.authPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewApigeeAdapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApigeeAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApigeeAdapter_Auth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/auth", func(c *gin.Context) {
		c.String(200, "Hello")
	})
	router.GET("/404", func(c *gin.Context) {
		c.String(404, "Hello")
	})
	router.POST("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"access_token": "token",
			"issued_at":    "0",
			"expires_in":   "1",
		})
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	type fields struct {
		endpoint     string
		clientID     string
		clientSecret string
		environment  string
		authPath     string
		httpClient   *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domain.TokenClaims
		wantErr bool
	}{
		{
			"invalid route returns error",
			fields{
				endpoint:   ts.URL,
				authPath:   "http://",
				httpClient: &http.Client{},
			},
			nil,
			true,
		},
		{
			"invalid status returns error",
			fields{
				endpoint:   ts.URL,
				authPath:   "/404",
				httpClient: &http.Client{},
			},
			nil,
			true,
		},
		{
			"runs correctly",
			fields{
				endpoint:   ts.URL,
				authPath:   "/auth",
				httpClient: &http.Client{},
			},
			&domain.TokenClaims{
				AccessToken: "token",
				IssuedAt:    "0",
				ExpiresIn:   "1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ApigeeAdapter{
				endpoint:     tt.fields.endpoint,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				environment:  tt.fields.environment,
				authPath:     tt.fields.authPath,
				httpClient:   tt.fields.httpClient,
			}
			got, err := a.Auth()
			if (err != nil) != tt.wantErr {
				t.Errorf("ApigeeAdapter.Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApigeeAdapter.Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
