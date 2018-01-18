package api

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/xwp/go-tide/src/tide"
)

// apiStub is a mock server for testing API requests.
var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	endpoint := r.URL.String()

	if endpoint == "/auth" {

		// Authentication responses.

		key := r.FormValue("api_key")
		secret := r.FormValue("api_secret")

		if key != "successId" || secret != "successSecret" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, `{}`)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{ "access_token": "verysecrettoken" }`)
		}
	} else {

		// Other responses.

		authKey := r.Header.Get("Authorization")
		authenticated := authKey != ""

		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{plugin:info}")
		case "POST":
			if authenticated {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, "{success:true}")
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintln(w, `{}`)
			}
		}
	}
}))

func TestClient_Authenticate(t *testing.T) {

	mockEndpoint := apiStub.URL + "/auth"

	type args struct {
		clientId     string
		clientSecret string
		authEndpoint string
	}

	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			c:    &Client{},
			args: args{
				clientId:     "successId",
				clientSecret: "successSecret",
				authEndpoint: mockEndpoint,
			},
			wantErr: false,
		},
		{
			name: "Fail",
			c:    &Client{},
			args: args{
				clientId:     "failId",
				clientSecret: "failSecret",
				authEndpoint: mockEndpoint,
			},
			wantErr: true,
		},
		{
			name: "Error",
			c:    &Client{},
			args: args{
				clientId:     "error",
				authEndpoint: "http://error",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Authenticate(tt.args.clientId, tt.args.clientSecret, tt.args.authEndpoint); (err != nil) != tt.wantErr {
				t.Errorf("Client.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendPayload(t *testing.T) {

	mockEndpoint := apiStub.URL + "/tide/v2/audit"

	type args struct {
		method   string
		endpoint string
		data     string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Unauthenticated GET",
			c:    &Client{},
			args: args{
				method:   "GET",
				endpoint: mockEndpoint,
				data:     "",
			},
			want:    "{plugin:info}",
			wantErr: false,
		},
		{
			name: "Authenticated GET",
			c:    &Client{
				&tide.Auth{
					AccessToken: "verysecrettoken",
				},
			},
			args: args{
				method:   "GET",
				endpoint: mockEndpoint,
				data:     "",
			},
			want:    "{plugin:info}",
			wantErr: false,
		},
		{
			name: "Unauthenticated POST",
			c:    &Client{},
			args: args{
				method:   "POST",
				endpoint: mockEndpoint,
				data:     "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Authenticated POST",
			c:    &Client{
				&tide.Auth{
					AccessToken: "verysecrettoken",
				},
			},
			args: args{
				method:   "POST",
				endpoint: mockEndpoint,
				data:     "",
			},
			want:    "{success:true}",
			wantErr: false,
		},
		{
			name: "Error",
			c:    &Client{},
			args: args{
				method:   "POST",
				endpoint: "http://error",
				data:     "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.SendPayload(tt.args.method, tt.args.endpoint, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SendPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.SendPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
