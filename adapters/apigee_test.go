package adapters

import (
	"io"
	"net/http"
	"reflect"
	"testing"
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
				endpoint:   ":#:",
				httpClient: &http.Client{},
			},
			args{
				header: &http.Header{},
			},
			0,
			nil,
			true,
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
