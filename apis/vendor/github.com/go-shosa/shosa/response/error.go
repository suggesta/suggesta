package response

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-shosa/shosa/log"
	"github.com/labstack/echo"
)

// ErrorData defines the information for error response messages and status.
type ErrorData struct {
	HTTPStatus       int    `json:"status",xml:"status"`
	Code             int    `json:"code,omitempty",xml:"code,omitempty"`
	Property         string `json:"property,omitempty",xml:"property,omitempty"`
	Message          string `json:"message",xml:"message"`
	DeveloperMessage string `json:"developerMessage",xml:"developerMessage"`
	MoreInfo         string `json:"moreInfo,omitempty",xml:"moreInfo,omitempty"`
}

// ErrorResponse defines the error response.
type ErrorResponse struct {
	Context          echo.Context
	ErrorInformation ErrorData
}

var (
	// BadRequest defines ErrorData data when returns 400 error response.
	BadRequest = ErrorData{
		HTTPStatus:       400,
		Property:         "bad_request",
		Message:          "The request is invalid or improperly formed.",
		DeveloperMessage: "The request is invalid or improperly formed.",
	}
	// Unauthorized defines ErrorData data when returns 400 error response.
	Unauthorized = ErrorData{
		HTTPStatus:       401,
		Property:         "unauthorized",
		Message:          "Invalid Credentials.",
		DeveloperMessage: "Invalid Credentials.",
	}
	// InsufficientPermission defines ErrorData data when returns 403 error response.
	InsufficientPermission = ErrorData{
		HTTPStatus:       403,
		Property:         "forbidden",
		Message:          "Does not have sufficient permissions.",
		DeveloperMessage: "Does not have sufficient permissions.",
	}
	// Forbidden defines ErrorData data when returns 403 error response.
	Forbidden = ErrorData{
		HTTPStatus:       403,
		Property:         "forbidden",
		Message:          "The request operation is forbidden and cannot be completed.",
		DeveloperMessage: "The request operation is forbidden and cannot be completed.",
	}
	// EndpointNotExists defines ErrorData data when is requested a invalid endpoint.
	EndpointNotExists = ErrorData{
		HTTPStatus:       404,
		Property:         "not_found",
		Message:          "Endpoint does not exist.",
		DeveloperMessage: "Endpoint does not exist.",
	}
	// DataNotFound defines ErrorData data when returns 404 error response.
	DataNotFound = ErrorData{
		HTTPStatus:       404,
		Property:         "not_found",
		Message:          "The requested data does not exist.",
		DeveloperMessage: "The requested data does not exist.",
	}
	// RequestEntityTooLarge defines ErrorData data when returns 413 error response.
	RequestEntityTooLarge = ErrorData{
		HTTPStatus:       413,
		Property:         "payload_too_large",
		Message:          "The data sent in the request is too large.",
		DeveloperMessage: "The data sent in the request is too large.",
	}
	// ResponseTooLarge defines ErrorData data when returns 413 error response.
	ResponseTooLarge = ErrorData{
		HTTPStatus:       413,
		Property:         "response_too_large",
		Message:          "The requested resource is too large to return.",
		DeveloperMessage: "The requested resource is too large to return.",
	}
	// RateLimitExceeded defines ErrorData data when returns 429 error response.
	RateLimitExceeded = ErrorData{
		HTTPStatus:       429,
		Property:         "rate_limit_exceeded",
		Message:          "API quota exceeded.",
		DeveloperMessage: "API quota exceeded.",
	}
	// DatabaseConnectError defines ErrorData data
	// when returns error response to fail to connect database.
	DatabaseConnectError = ErrorData{
		HTTPStatus:       500,
		Property:         "internal_server_error",
		Message:          "Failed to refer data storage.",
		DeveloperMessage: "Failed to refer data storage.",
	}
	// InternalServerError defines ErrorData data when returns 500 error response.
	InternalServerError = ErrorData{
		HTTPStatus:       500,
		Property:         "internal_server_error",
		Message:          "The request failed due to an internal error.",
		DeveloperMessage: "The request failed due to an internal error.",
	}
	// ServiceUnavailable defines ErrorData data when returns error response during matainance.
	ServiceUnavailable = ErrorData{
		HTTPStatus:       503,
		Property:         "service_unavailable",
		Message:          "Overloaded with requests. Please try again later..",
		DeveloperMessage: "Overloaded with requests. Please try again later..",
	}
)

// NewErrorResponse creates a new ErrorResponse.
func NewErrorResponse(c echo.Context, e ErrorData) *ErrorResponse {
	return &ErrorResponse{
		Context:          c,
		ErrorInformation: e,
	}
}

// DecodeErrorData create a new ErrorData from json text.
func DecodeErrorData(jsonText string) (d ErrorData, err error) {
	dec := json.NewDecoder(strings.NewReader(jsonText))
	err = dec.Decode(&d)

	return d, err
}

// DecodeResopnseAsErrorData create a new ErrorData from internal API response.
// when invalid JSON response catches, InternalServerError is created.
func DecodeResopnseAsErrorData(body string) (d ErrorData) {
	d, err := DecodeErrorData(body)

	if err != nil {
		d = InternalServerError
		d.DeveloperMessage = body
	}
	return d
}

// Error returns error message string.
func (er *ErrorResponse) Error() string {
	return er.ErrorInformation.DeveloperMessage
}

// JSON sends a JSON response with status code.
func (er *ErrorResponse) JSON() (err error) {
	ei := er.ErrorInformation
	log.Error(er.Error())
	return er.Context.JSON(ei.HTTPStatus, ei)
}

// XML sends a XML response with status code.
func (er *ErrorResponse) XML() (err error) {
	ei := er.ErrorInformation
	log.Error(er.Error())
	return er.Context.XML(ei.HTTPStatus, ei)
}

// String sends a String response with status code.
func (er *ErrorResponse) String() (err error) {
	ei := er.ErrorInformation
	str := fmt.Sprintf("status: %d, code: %d, property: %s, message: %s, develop: %s, more: %s", ei.HTTPStatus, ei.Code, ei.Property, ei.Message, ei.DeveloperMessage, ei.MoreInfo)
	log.Error(er.Error())
	return er.Context.String(ei.HTTPStatus, str)
}
