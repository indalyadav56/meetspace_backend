package utils

import (
	"net/http"
	"time"
)

// "extensions": {
//  "user_session": "xyz"
// }
// "refresh_in": 1800, //seconds
// "refresh_token": "xyz123"
// "pagination": {
//   "total": 50,
//   "offset": 0,
//   "limit": 10
// }
// {
//   "status_code": 200,
//   "message": "Successfully logged in!",
//   "data": {
//     "first_name": "Indal"
//   },
//   "error": null,
//   "request_id": "",
//   "timestamp": "2024-01-01T12:00:00Z",
//   "metadata": {
//     "api_version": "v1",
//     "server_info": {
//   "server_type": "nginx",
//   "server_version": "1.21.3",
//   "environment": "production",
//   "hostname": "example.com"
// }
//   }
// }
// "debug": {
//     "stack_trace": xxx,
//     "request_params": xxx
//   }
// "messages": {
//   "en": "Logged in",
//   "es": "Sesión iniciada"
// }
// Human readable status/error codes:
// Human readable status/error codes:
// "status": "SUCCESS"
// "error_code": "AUTH_ERROR"
//


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
func SuccessResponse(message string, data interface{}) Response {
    metadata := Metadata{
        ApiVersion: "v1",
        ServerInfo: ServerInfo{
            Hostname: "localhost",
        },
    }
    return Response{
        Status: "success",
        StatusCode: http.StatusOK,
        Message: message,
        Data:    data,
        Metadata: metadata,
    }
}

// ErrorResponse creates an error response with an optional message and errors
func ErrorResponse(message string, errorData interface{}) Response {
    return Response{
        Status: "error",
        StatusCode: http.StatusBadRequest,
        Message: message,
        Error: errorData,
    }
}

// ErrorResponse creates an error response with an optional message and errors
func NotFoundErrorResponse(message string, errorData interface{}) Response {
    return Response{
        Status: "error",
        StatusCode: http.StatusNotFound,
        Message: message,
        Error: errorData,
    }
}


func InternalError(message string, data []interface{}) Response {
    return Response{
      StatusCode: http.StatusInternalServerError,
      Message: message, 
      Error: data,
    } 
}
