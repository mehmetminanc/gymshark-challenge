package handlers

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_validateRequest(t *testing.T) {
	type args struct {
		req *CalculateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "should return success when order is negative",
			args: args{
				req: &CalculateRequest{
					Order: 501,
					Sizes: []int{250, 500, 1000},
				},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: false,
		},
		{
			name: "should return 400 when a pack size is negative",
			args: args{
				req: &CalculateRequest{
					Order: 501,
					Sizes: []int{250, -500, 1000},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			},
			wantErr: true,
		},
		{
			name: "should return 400 when order is negative",
			args: args{
				req: &CalculateRequest{
					Order: -13,
					Sizes: []int{},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_New(t *testing.T) {
	handler := New()
	tests := []struct {
		name    string
		body    string
		want    string
		wantErr bool
	}{
		{
			name:    "default size",
			body:    `{"order": 1002, "sizes": []}`,
			want:    `{"1000":1, "250":1}`,
			wantErr: false,
		},
		{
			name:    "custom size",
			body:    `{"order": 1002, "sizes": [1,2,5,10,20,50,100,200]}`,
			want:    `{"200":5, "2":1}`,
			wantErr: false,
		},
		{
			name:    "malformed request",
			body:    `{"order": 1002, "si`,
			want:    ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := handler(context.Background(), events.APIGatewayProxyRequest{
				Body: tt.body,
			})
			if err != nil {
				if tt.wantErr {
					assert.Equal(t, http.StatusBadRequest, response.StatusCode, "should return StatusBadRequest")
				} else {
					t.Errorf("validateRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else {
				assert.JSONEq(t, tt.want, response.Body)
			}
		})
	}
}