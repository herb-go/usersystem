package usersystem

import "context"

type ContextKey string

var ContextKeyUsersystem = ContextKey("usersystem")
var ContextKeyUID = ContextKey("uid")

func UIDContext(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, ContextKeyUID, uid)
}
func GetUID(ctx context.Context) string {
	return ctx.Value(ContextKeyUID).(string)
}
