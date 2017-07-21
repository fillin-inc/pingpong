package pingpong

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type expected struct {
	HTTPStatus  int
	ContentType string
	Body        string
}

func TestHandler(t *testing.T) {
	e := expected{
		HTTPStatus:  http.StatusOK,
		ContentType: "text/plain",
		Body:        PONG,
	}

	ts := httptest.NewServer(http.HandlerFunc(Handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != e.HTTPStatus {
		t.Errorf(
			"Handler return unexpected HTTP Status(expected:%d actual:%d)",
			e.HTTPStatus,
			res.StatusCode,
		)
	}

	if res.Header.Get("Content-Type") != e.ContentType {
		t.Errorf(
			"Handler return unexpected Content-Type(expected:%s actual:%s)",
			e.ContentType,
			res.Header.Get("Content-Type"),
		)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != e.Body {
		t.Errorf(
			"Handler return unexpected Body(expected:%s actual:%s)",
			e.Body,
			string(body),
		)
	}
}
