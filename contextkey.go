package usersystem

import "context"

type ContextKey string

var ContextKeyUsersystem = ContextKey("usersystem")

func GetUsersystem(ctx context.Context) *UserSystem {
	return ctx.Value(ContextKeyUsersystem).(*UserSystem)
}

var ContextKeyUID = ContextKey("uid")

func UIDContext(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, ContextKeyUID, uid)
}
func GetUID(ctx context.Context) string {
	return ctx.Value(ContextKeyUID).(string)
}

var ContextKeySession = ContextKey("session")

func SessionContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, ContextKeySession, session)
}

func GetSession(ctx context.Context) *Session {
	return ctx.Value(ContextKeySession).(*Session)
}
