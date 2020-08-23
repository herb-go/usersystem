package httpusersystem

import (
	"context"
	"net/http"

	"github.com/herb-go/usersystem"
)

var ContextKeyRequest = usersystem.ContextKey("http/request")

func RequestContext(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, ContextKeyRequest, req)
}

func GetRequest(ctx context.Context) *http.Request {
	v := ctx.Value(ContextKeyRequest)
	r, ok := v.(*http.Request)
	if !ok {
		return nil
	}
	return r
}
