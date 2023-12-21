package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	middlewareChi "github.com/go-chi/chi/v5/middleware"
)

type responseLogger struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (r *responseLogger) Write(b []byte) (int, error) {
	r.buf.Write(b)
	return r.ResponseWriter.Write(b)
}

func (m *Middleware) Logger(next http.Handler) http.Handler {
	m.log.Debug("logger middlewares enabled")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		args := []any{
			"request_method", r.Method,
			"request_path", r.URL.Path,
			"request_user_agent", r.UserAgent(),
			"request_id", middlewareChi.GetReqID(r.Context()),
		}

		if r.Body != nil {
			body, err := io.ReadAll(r.Body)
			if err == nil {
				args = append(args, "request_body", string(body))
				r.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		buf := bytes.NewBuffer(nil)
		rl := &responseLogger{w, buf}
		ww := middlewareChi.NewWrapResponseWriter(rl, r.ProtoMajor)
		timeStart := time.Now()

		defer func() {
			args = append(args, []any{
				"response_status", ww.Status(),
				"response_bytes", ww.BytesWritten(),
				"response_duration", time.Since(timeStart).String(),
			}...)

			if ww.Status() >= http.StatusOK && ww.Status() < http.StatusBadRequest {
				responseBody := buf.String()
				args = append(args, "response_body", responseBody)
			}

			m.log.Infow("request", args...)
		}()

		next.ServeHTTP(ww, r)
	})
}
