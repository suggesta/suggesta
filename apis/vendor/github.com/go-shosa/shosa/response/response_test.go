package response

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

type (
	user struct {
		ID   int    `json:"id" xml:"id" form:"id"`
		Name string `json:"name" xml:"name" form:"name"`
	}
)

const (
	userJSON = `{"id":1,"name":"Jon Snow"}`
	userXML  = `<user><id>1</id><name>Jon Snow</name></user>`
)

func TestNew(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	status := 200
	item := "hello, human!!"
	expected := &Response{
		Context:    c,
		HTTPStatus: status,
		Item:       item,
	}

	actual := New(c, status, item)
	if expected.HTTPStatus != actual.HTTPStatus {
		t.Fatalf("actual.HTTPStatus should be equal expected.HTTPStatus. expected:%d, actual:%d", expected.HTTPStatus, actual.HTTPStatus)
	}
	if expected.Item != actual.Item {
		t.Fatalf("actual.Item should be equal expected.Item. expected:%v, actual:%v", expected.Item, actual.Item)
	}
}

func TestJSON(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	resp := New(c, http.StatusOK, user{1, "Jon Snow"})
	err := resp.JSON()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Code {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Code)
	}
	if echo.MIMEApplicationJSONCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %s, but actual %s", echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if userJSON != rec.Body.String() {
		t.Errorf("Response body is expected %s, but actual %s", userJSON, rec.Body.String())
	}
}

func TestXML(t *testing.T) {
	actual := xml.Header + userXML

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	resp := New(c, http.StatusOK, user{1, "Jon Snow"})
	err := resp.XML()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Code {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Code)
	}
	if echo.MIMEApplicationXMLCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %s, but actual %s", echo.MIMEApplicationXMLCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if actual != rec.Body.String() {
		t.Errorf("Response body is expected %s, but actual %s", actual, rec.Body.String())
	}
}

func TestString(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	resp := New(c, http.StatusOK, "Hello, World!")
	err := resp.String()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Code {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Code)
	}
	if echo.MIMETextPlainCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %s, but actual %s", echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if "Hello, World!" != rec.Body.String() {
		t.Errorf("Response body is expected %s, but actual %s", "Hello, World!", rec.Body.String())
	}
}
