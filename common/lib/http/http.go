package http

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var contextKeyForRequestID interface{}
var (
	httpInstance     *HttpClientImpl
	httpOnce         sync.Once
	globalRequestKey interface{}
)

type HttpConfig struct {
	Timeout      time.Duration
	RequestIDKey interface{}
}

type HttpClient interface {
	Get(ctx context.Context, url string, headers map[string]string) ([]byte, error)
	Post(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, error)
	Patch(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, error)
}

type HttpClientImpl struct {
	client       *http.Client
	requestIDKey interface{}
}

// ---- Initialization ----

func Init(cfg HttpConfig) *HttpClientImpl {
	httpOnce.Do(func() {
		if cfg.Timeout == 0 {
			cfg.Timeout = 15 * time.Second
		}

		httpInstance = &HttpClientImpl{
			client:       &http.Client{Timeout: cfg.Timeout},
			requestIDKey: cfg.RequestIDKey,
		}
		contextKeyForRequestID = cfg.RequestIDKey
		globalRequestKey = cfg.RequestIDKey
	})
	return httpInstance
}

// Get returns the singleton instance (auto-init with default if needed)
func Get() *HttpClientImpl {
	if httpInstance == nil {
		Init(HttpConfig{
			Timeout:      15 * time.Second,
			RequestIDKey: "tracing_id",
		})
	}
	return httpInstance
}

// ---- Internal helpers ----

func (h *HttpClientImpl) buildRequest(ctx context.Context, method, url string, headers map[string]string, body []byte) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		log.Printf("[HTTP] Failed to create %s request: %v", method, err)
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	
	// Add X-Request-ID if present in context
	if ctx != nil && h.requestIDKey != nil {
		if rid := ctx.Value(h.requestIDKey); rid != nil {
			if s, ok := rid.(string); ok && s != "" {
				req.Header.Set("tracing_id", s)
			}
		}
	}

	return req, nil
}

func (h *HttpClientImpl) do(ctx context.Context, req *http.Request) ([]byte, error) {
	start := time.Now()
	resp, err := h.client.Do(req)
	duration := time.Since(start)

	if err != nil {
		log.Printf("[HTTP] %s %s failed: %v", req.Method, req.URL.String(), err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[HTTP] Failed to read response: %v", err)
		return nil, err
	}

	log.Printf("[HTTP] %s %s -> %d (%dms)", req.Method, req.URL.String(), resp.StatusCode, duration.Milliseconds())
	return body, nil
}

// ---- Public methods ----

func (h *HttpClientImpl) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	req, err := h.buildRequest(ctx, http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return h.do(ctx, req)
}

func (h *HttpClientImpl) Post(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := h.buildRequest(ctx, http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}
	return h.do(ctx, req)
}

func (h *HttpClientImpl) Patch(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := h.buildRequest(ctx, http.MethodPatch, url, headers, body)
	if err != nil {
		return nil, err
	}
	return h.do(ctx, req)
}

// Optional: expose global request key for other packages
func GetRequestIDKey() interface{} {
	return globalRequestKey
}
