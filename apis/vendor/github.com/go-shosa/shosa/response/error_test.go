package response

import (
	"fmt"
	"testing"
)

func TestDecodeResopnseAsErrorDataValid(t *testing.T) {
	httpStatus := 400
	code := 400
	property := "bad_request"
	message := "message."
	developerMessage := "developerMessage."
	moreInfo := "moreInfo."
	body := fmt.Sprintf("{\"status\":%d,\"code\":%d,\"property\":\"%s\",\"message\":\"%s\",\"developerMessage\":\"%s\",\"moreInfo\":\"%s\"}", httpStatus, code, property, message, developerMessage, moreInfo)

	result := DecodeResopnseAsErrorData(body)

	if result.HTTPStatus != httpStatus {
		t.Errorf("Response http status is expected %d, but actual %d", httpStatus, result.HTTPStatus)
	}
	if result.Code != code {
		t.Errorf("Response http code is expected %d, but actual %d", code, result.Code)
	}
	if result.Property != property {
		t.Errorf("Response http property is expected %s, but actual %s", property, result.Property)
	}
	if result.Message != message {
		t.Errorf("Response http message is expected %s, but actual %s", message, result.Message)
	}
	if result.DeveloperMessage != developerMessage {
		t.Errorf("Response http developerMessage is expected %s, but actual %s", developerMessage, result.DeveloperMessage)
	}
	if result.MoreInfo != moreInfo {
		t.Errorf("Response http moreInfo is expected %s, but actual %s", moreInfo, result.MoreInfo)
	}
}

func TestDecodeResopnseAsErrorDataInvalid(t *testing.T) {
	body := "invalid JSON body"

	result := DecodeResopnseAsErrorData(body)

	if result.HTTPStatus != InternalServerError.HTTPStatus {
		t.Errorf("Response http status is expected %d, but actual %d", InternalServerError.HTTPStatus, result.HTTPStatus)
	}
	if result.Property != InternalServerError.Property {
		t.Errorf("Response http property is expected %s, but actual %s", InternalServerError.Property, result.Property)
	}
	if result.Message != InternalServerError.Message {
		t.Errorf("Response http message is expected %s, but actual %s", InternalServerError.Property, result.Message)
	}
	if result.DeveloperMessage != body {
		t.Errorf("Response http developerMessage is expected %s, but actual %s", InternalServerError.Property, body)
	}
}
