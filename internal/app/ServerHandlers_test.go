package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type want struct {
	contentType string
	statusCode  int
	regexpLink  string
}

type request struct {
	method string
	url    string
	body   string
}

type tstRequest struct {
	nameTest string
	request  request
	want     want
}

func TestGetShortUrl(t *testing.T) {

	tests := []tstRequest{
		{

			nameTest: "standard test",
			request: request{
				method: http.MethodPost,
				url:    "/",
				body:   "www.google.com",
			},
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
				regexpLink:  "e",
			}},
	}

	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			request := httptest.NewRequest(tt.request.method, tt.request.url, strings.NewReader(tt.request.body))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(URLHandler)
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := ioutil.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)
			fmt.Println(userResult)

			//var user User
			//err = json.Unmarshal(userResult, &user)
			//require.NoError(t, err)
			//
			//assert.Equal(t, tt.want.user, user)
		})
	}
}

//
//func TestURLHandler(t *testing.T) {
//	type args struct {
//		writer  http.ResponseWriter
//		request *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			URLHandler(tt.args.writer, tt.args.request)
//		})
//	}
//}
//
//func Test_getFullURL(t *testing.T) {
//	type args struct {
//		writer  http.ResponseWriter
//		request *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			getFullURL(tt.args.writer, tt.args.request)
//		})
//	}
//}
