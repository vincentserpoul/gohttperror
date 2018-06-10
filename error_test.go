package gohttperror

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrResponse_Render(t *testing.T) {

	tests := []struct {
		name        string
		errResponse *ErrResponse
		wantErr     bool
	}{
		{
			name:        "error response",
			errResponse: &ErrResponse{},
			wantErr:     false,
		},
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", ``, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.errResponse.Render(w, r); (err != nil) != tt.wantErr {
				t.Errorf(
					"ErrResponse.Render() error = %v, wantErr %v",
					err, tt.wantErr)
			}
		})
	}
}

func TestErrBadRequest(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name            string
		args            args
		wantedLogFields []string
	}{
		{
			name: "working",
			args: args{
				err: fmt.Errorf("Pacho"),
			},
			wantedLogFields: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrBadRequest(tt.args.err)
			if got.HTTPStatusCode != http.StatusBadRequest {
				t.Errorf(
					"ErrBadRequest() doesn't return the right status"+
						" got %d instead of %d",
					got.HTTPStatusCode,
					http.StatusBadRequest,
				)
				return
			}
		})
	}
}

func TestErrInternal(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name            string
		args            args
		wantedLogFields []string
	}{
		{
			name: "working",
			args: args{
				err: fmt.Errorf("Pacho"),
			},
			wantedLogFields: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrInternal(tt.args.err)
			if got.HTTPStatusCode != http.StatusInternalServerError {
				t.Errorf(
					"ErrInternal() doesn't return the right status"+
						" got %d instead of %d",
					got.HTTPStatusCode,
					http.StatusInternalServerError,
				)
				return
			}
		})
	}
}

func TestErrForbidden(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name            string
		args            args
		wantedLogFields []string
	}{
		{
			name: "working",
			args: args{
				err: fmt.Errorf("Pacho"),
			},
			wantedLogFields: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrForbidden(
				tt.args.err,
			)
			if got.HTTPStatusCode != http.StatusForbidden {
				t.Errorf(
					"ErrForbidden() doesn't return the right status"+
						" got %d instead of %d",
					got.HTTPStatusCode,
					http.StatusForbidden,
				)
				return
			}
		})
	}
}

func TestErrUnauthorized(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name            string
		args            args
		wantedLogFields []string
	}{
		{
			name: "working",
			args: args{
				err: fmt.Errorf("Pacho"),
			},
			wantedLogFields: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrUnauthorized(tt.args.err)
			if got.HTTPStatusCode != http.StatusUnauthorized {
				t.Errorf(
					"ErrUnauthorized() doesn't return the right status"+
						" got %d instead of %d",
					got.HTTPStatusCode,
					http.StatusUnauthorized,
				)
				return
			}
		})
	}
}
