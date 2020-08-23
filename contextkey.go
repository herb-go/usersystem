package usersystem

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
)

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

var ContextKeySessionType = ContextKey("sessiontype")

func SessionTypeContext(ctx context.Context, st SessionType) context.Context {
	return context.WithValue(ctx, ContextKeySessionType, st)
}
func GetSessionType(ctx context.Context) SessionType {
	return ctx.Value(ContextKeySessionType).(SessionType)
}

var ContextKeyPayloads = ContextKey("usersession.payloads")

func PayloadsContext(ctx context.Context, payloads *authority.Payloads) context.Context {
	return context.WithValue(ctx, ContextKeyPayloads, payloads)
}

func GetPayloads(ctx context.Context) *authority.Payloads {
	return ctx.Value(ContextKeyPayloads).(*authority.Payloads)
}
