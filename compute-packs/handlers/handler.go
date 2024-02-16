package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/aws/aws-lambda-go/events"

	"github.com/mehmetminanc/gymshark-challenge/compute-packs/internal/algo"
)

var (
	defaultSizes = []int{250, 500, 1000, 2000, 5000}
)

type CalculateRequest struct {
	Order int   `json:"order,omitempty"`
	Sizes []int `json:"sizes,omitempty"`
}

func New() func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// similar to http middlewares, initialize dependencies here and use in
	// returned handler function.
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var req CalculateRequest
		err := json.Unmarshal([]byte(request.Body), &req)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			}, err
		}

		if req.Sizes == nil || len(req.Sizes) == 0 {
			req.Sizes = defaultSizes
		}

		response, err2 := validateRequest(req)
		if err2 != nil {
			return response, err2
		}

		// Execute
		packing := algo.ComputePacking(req.Sizes, req.Order)
		resp, err := json.MarshalIndent(packing, "", "  ")
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			}, err
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type":                 "text/json",
				"Access-Control-Allow-Headers": "Content-Type",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
				"X-Content-Type-Options":       "nosniff",
			},
			Body: string(resp),
		}, nil
	}
}

func validateRequest(req CalculateRequest) (events.APIGatewayProxyResponse, error) {
	// Validate
	if req.Order <= 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("order can't be 0 or below: %d", req.Order)
	}
	if slices.Min(req.Sizes) <= 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("no pack size can be 0 or below: %d", req.Order)
	}
	return events.APIGatewayProxyResponse{}, nil
}
