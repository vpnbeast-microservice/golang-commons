package commons

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type DummyServiceOptions struct {
	Foo    string `env:"FOO"`
	Bar    string `env:"BAR"`
	Length int    `env:"LENGTH"`
}

func TestInitOptions(t *testing.T) {
	_ = os.Setenv("ACTIVE_PROFILE", "unit-test")
	_ = os.Setenv("CONFIG_PATH", "./")
	opts := &DummyServiceOptions{}
	err := InitOptions(opts, "dummy-service")
	assert.Nil(t, err)
	assert.NotNil(t, opts.Foo)
	assert.NotNil(t, opts.Bar)
	assert.Equal(t, "foo", opts.Foo)
	assert.Equal(t, "bar", opts.Bar)
	assert.Equal(t, 12, opts.Length)
}

func TestInitOptionsRemoteServer(t *testing.T) {
	appName := "dummy-service"
	profile := "remote"
	_ = os.Setenv("ACTIVE_PROFILE", profile)
	_ = os.Setenv("CONFIG_PATH", "./")

	expectedResponse := "dummy-service:\n  foo: foo\n  bar: bar\n  length: 12"
	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Path = fmt.Sprintf("%s-%s.yaml", appName, profile)
		writer.WriteHeader(200)
		if _, err := fmt.Fprint(writer, expectedResponse); err != nil {
			t.Fatalf("a fatal error occured while writing response body: %s", err.Error())
		}
	}))
	defer mockServer.Close()

	_ = os.Setenv("CONFIG_SERVICE_HOST", "127.0.0.1")
	_ = os.Setenv("CONFIG_SERVICE_PORT", strings.Split(mockServer.Listener.Addr().String(), ":")[1])

	opts := &DummyServiceOptions{}
	err := InitOptions(opts, appName)
	assert.Nil(t, err)
	assert.NotNil(t, opts.Foo)
	assert.NotNil(t, opts.Bar)
	assert.Equal(t, "foo", opts.Foo)
	assert.Equal(t, "bar", opts.Bar)
	assert.Equal(t, 12, opts.Length)
}
