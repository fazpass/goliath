package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
)

type (
	RecovererOptions struct {
		Debug    bool
		Sentry   bool
		Response RecovererResponse
	}

	RecovererResponse struct {
		Status      int
		ContentType string
		Body        interface{}
	}
)

func Recoverer(options RecovererOptions) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					if options.Debug {
						debug.PrintStack()
					}

					if options.Sentry {
						errStr := fmt.Sprint(rvr)
						sentry.CaptureException(errors.New(errStr))
					}

					w = CreateResponse(w, options.Response.ContentType, options.Response.Status, options.Response.Body)
				}
			}()

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)

	}

	return f
}

func CreateResponse(w http.ResponseWriter, contentType string, status int, body interface{}) http.ResponseWriter {
	if contentType == "" {
		contentType = "application/json"
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)

	if contentType == "application/json" {
		jData, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		w.Write(jData)
	} else {
		w.Write([]byte(body.(string)))
	}

	return w
}
