package utils

import (
	"net/http"
	"time"
)


type ServerInfo struct {
    Hostname string `json:"hostname"`
}

type Metadata struct {
    ApiVersion string `json:"api_version" default:"v1"`
    ServerInfo ServerInfo `json:"server_info"`
}


// Response represents a generic API response structure
type Response struct {
    StatusCode int         `json:"status_code"`
    Status string         `json:"status"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data"`
    Error       interface{} `json:"error"`
    RequestId   string `json:"request_id"`
    TimeStamp  time.Time   `json:"time_stamp"`
    Metadata Metadata `json:"metadata"`
}

// SuccessResponse creates a success response with optional message and data
func SuccessResponse(message string, data interface{}) *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
        Status: "success",
        StatusCode: http.StatusOK,
        Message: message,
        Data:    data,
        Metadata: metadata,
    }
}

// ErrorResponse creates an error response with an optional message and errors
func ErrorResponse(message string, errorData interface{}) *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
        Status: "error",
        StatusCode: http.StatusBadRequest,
        Message: message,
        Error: errorData,
        Metadata: metadata,
    }
}

func UnauthorizedResponse(message string) *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
        Status: "error",
        StatusCode: http.StatusUnauthorized,
        Message: message,
        Metadata: metadata,
    }
}

// ErrorResponse creates an error response with an optional message and errors
func NotFoundErrorResponse(message string, errorData interface{}) *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
        Status: "error",
        StatusCode: http.StatusNotFound,
        Message: message,
        Error: errorData,
        Metadata: metadata,
    }
}

func InternalError(message string, data []interface{}) *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
      StatusCode: http.StatusInternalServerError,
      Message: message, 
      Error: data,
      Metadata: metadata,
    } 
}

// NoContentResponse creates a 204 No Content response
func NoContentResponse() *Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return &Response{
        StatusCode: http.StatusNoContent,
        Metadata: metadata,
    }
}