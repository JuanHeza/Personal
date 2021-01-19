package controllers

import (
	"net/http"
	"testing"
)

func TestUpdateStatics(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStatics(tt.args.w, tt.args.r)
		})
	}
}
