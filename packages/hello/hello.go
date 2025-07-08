package main

import (
    "context"
    "encoding/json"
    "fmt"
)

type Response struct {
    StatusCode int               `json:"statusCode"`
    Body       string           `json:"body"`
    Headers    map[string]string `json:"headers,omitempty"`
}

type Request struct {
    HTTPMethod string            `json:"httpMethod"`
    Path       string            `json:"path"`
    Headers    map[string]string `json:"headers"`
    Body       string            `json:"body"`
    QueryStringParameters map[string]string `json:"queryStringParameters"`
}

func Main(ctx context.Context, in Request) (*Response, error) {
    // Handle different HTTP methods
    switch in.HTTPMethod {
    case "GET":
        return handleGet(in)
    case "POST":
        return handlePost(in)
    default:
        return &Response{
            StatusCode: 405,
            Body:       `{"error": "Method not allowed"}`,
            Headers: map[string]string{
                "Content-Type": "application/json",
            },
        }, nil
    }
}

func handleGet(in Request) (*Response, error) {
    name := in.QueryStringParameters["name"]
    if name == "" {
        name = "World"
    }
    
    response := map[string]interface{}{
        "message": fmt.Sprintf("Hello, %s!", name),
        "method":  "GET",
        "path":    in.Path,
    }
    
    body, _ := json.Marshal(response)
    
    return &Response{
        StatusCode: 200,
        Body:       string(body),
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    }, nil
}

func handlePost(in Request) (*Response, error) {
    var requestData map[string]interface{}
    if err := json.Unmarshal([]byte(in.Body), &requestData); err != nil {
        return &Response{
            StatusCode: 400,
            Body:       `{"error": "Invalid JSON"}`,
            Headers: map[string]string{
                "Content-Type": "application/json",
            },
        }, nil
    }
    
    response := map[string]interface{}{
        "message": "Data received successfully",
        "method":  "POST",
        "data":    requestData,
    }
    
    body, _ := json.Marshal(response)
    
    return &Response{
        StatusCode: 200,
        Body:       string(body),
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    }, nil
}

