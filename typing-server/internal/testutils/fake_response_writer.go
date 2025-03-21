package testutils

import (
	"bytes"
	"errors"
	"net/http"
)

// FakeResponseWriter はテスト用のhttp.ResponseWriterの実装です
type FakeResponseWriter struct {
	Header_     http.Header
	Body        *bytes.Buffer
	StatusCode  int
	FailOnWrite bool
	Wrote       bool
}

// NewFakeResponseWriter は新しいFakeResponseWriterインスタンスを作成します
func NewFakeResponseWriter() *FakeResponseWriter {
	return &FakeResponseWriter{
		Header_:     make(http.Header),
		Body:        new(bytes.Buffer),
		FailOnWrite: true,
	}
}

func (f *FakeResponseWriter) Header() http.Header {
	return f.Header_
}

func (f *FakeResponseWriter) Write(b []byte) (int, error) {
	if f.FailOnWrite && !f.Wrote {
		f.Wrote = true
		return 0, errors.New("simulated write error")
	}
	return f.Body.Write(b)
}

func (f *FakeResponseWriter) WriteHeader(statusCode int) {
	f.StatusCode = statusCode
}