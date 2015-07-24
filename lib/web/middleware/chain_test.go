package middleware

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jgroeneveld/gotrix/lib/web"
	"github.com/jgroeneveld/gotrix/lib/web/ctx"
)

func TestChain(t *testing.T) {
	callstack := &callstack{}

	chain := NewChain(
		csmw(callstack, "1"),
		csmw(callstack, "2"),
	)

	handler := func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		callstack.Called("handler")
		return nil
	}

	finalhandler := chain.Bind(handler)
	finalhandler(nil, nil, nil)

	if callstack.Join() != "1_before, 2_before, handler, 2_after, 1_after" {
		t.Errorf("callstack not as expected: %q", callstack.Join())
	}
}

func TestStack_NestedStacks(t *testing.T) {
	callstack := &callstack{}

	innerStack := NewChain(
		csmw(callstack, "inner1"),
		csmw(callstack, "inner2"),
	)

	outerStack := NewChain(
		csmw(callstack, "outer1"),
		csmw(callstack, "outer2"),
		innerStack,
	)

	handler := func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		callstack.Called("handler")
		return nil
	}

	finalhandler := outerStack.Bind(handler)
	finalhandler(nil, nil, nil)

	if callstack.Join() != "outer1_before, outer2_before, inner1_before, inner2_before, handler, inner2_after, inner1_after, outer2_after, outer1_after" {
		t.Errorf("callstack not as expected: %q", callstack.Join())
	}
}

type callstack struct {
	calls []string
}

func (s *callstack) Called(msg string) {
	s.calls = append(s.calls, msg)
}

func (s *callstack) Join() string {
	return strings.Join(s.calls, ", ")
}

func csmw(callstack *callstack, msg string) *callstackMiddleware {
	return &callstackMiddleware{
		callstack: callstack,
		Msg:       msg,
	}
}

type callstackMiddleware struct {
	callstack *callstack
	Msg       string
}

func (mw *callstackMiddleware) Bind(next web.HTTPHandle) web.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		mw.callstack.Called(mw.Msg + "_before")
		err := next(rw, r, c)
		mw.callstack.Called(mw.Msg + "_after")
		return err
	}
}
