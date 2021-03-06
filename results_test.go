package revel

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that the render response is as expected.
func TestBenchmarkRender(t *testing.T) {
	startFakeBookingApp()
	resp := httptest.NewRecorder()
	c := NewController(NewRequest(showRequest), NewResponse(resp), nil)
	c.SetAction("Hotels", "Show")
	result := Hotels{c}.Show(3)
	result.Apply(c.Request, c.Response)
	if !strings.Contains(resp.Body.String(), "300 Main St.") {
		t.Errorf("Failed to find hotel address in action response:\n%s", resp.Body)
	}
}

func BenchmarkRenderChunked(b *testing.B) {
	startFakeBookingApp()
	resp := httptest.NewRecorder()
	resp.Body = nil
	c := NewController(NewRequest(showRequest), NewResponse(resp), nil)
	c.SetAction("Hotels", "Show")
	Config.SetOption("results.chunked", "true")
	b.ResetTimer()

	hotels := Hotels{c}
	for i := 0; i < b.N; i++ {
		hotels.Show(3).Apply(c.Request, c.Response)
	}
}

func BenchmarkRenderNotChunked(b *testing.B) {
	startFakeBookingApp()
	resp := httptest.NewRecorder()
	resp.Body = nil
	c := NewController(NewRequest(showRequest), NewResponse(resp), nil)
	c.SetAction("Hotels", "Show")
	Config.SetOption("results.chunked", "false")
	b.ResetTimer()

	hotels := Hotels{c}
	for i := 0; i < b.N; i++ {
		hotels.Show(3).Apply(c.Request, c.Response)
	}
}
