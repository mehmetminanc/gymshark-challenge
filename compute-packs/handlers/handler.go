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

const (
	upperLimit = 100_000_000
)

var (
	defaultSizes = []int{250, 500, 1000, 2000, 5000}
	corsHeaders  = map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
		"X-Content-Type-Options":       "nosniff",
	}
)

// CalculateRequest input object
type CalculateRequest struct {
	Order int   `json:"order,omitempty"`
	Sizes []int `json:"sizes,omitempty"`
}

// validateRequest apply defaults and validate request
func (req *CalculateRequest) validateRequest() error {
	// Apply defaults
	if req.Sizes == nil || len(req.Sizes) == 0 {
		req.Sizes = defaultSizes
	}

	// Validate
	if req.Order <= 0 {
		return fmt.Errorf("order can't be 0 or below: %d", req.Order)
	}

	if req.Order > upperLimit { // Another 0 and we are out of mem.
		return fmt.Errorf("order can't be above %d, it was %d", upperLimit, req.Order)
	}

	if slices.Min(req.Sizes) <= 0 { // Everything shall be positive
		return fmt.Errorf("no pack size can be 0 or below: %d", req.Order)
	}
	return nil
}

type ErrorResponse struct {
	ErrorResponse string `json:"error"`
}

func New() func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// similar to http middlewares, initialize dependencies here and use in
	// returned handler function.
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var req CalculateRequest
		err := json.Unmarshal([]byte(request.Body), &req)
		if err != nil {
			return failure(err)
		}

		err = req.validateRequest()
		if err != nil {
			return failure(err)
		}

		// Execute
		packing := algo.ComputePacking(req.Sizes, req.Order)
		resp, _ := json.MarshalIndent(packing, "", "  ")

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    corsHeaders,
			Body:       string(resp),
		}, nil
	}
}

func failure(err error) (events.APIGatewayProxyResponse, error) {
	body, _ := json.MarshalIndent(ErrorResponse{
		ErrorResponse: err.Error(),
	}, "", "  ")

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(body),
		Headers:    corsHeaders,
	}, nil
}
