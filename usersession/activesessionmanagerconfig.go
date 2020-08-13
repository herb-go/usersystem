package usersession

import (
	"context"
	"time"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type Config struct {
	Supported bool
	Duration  time.Duration
}

var ContextKeyConfig = usersystem.ContextKey("usersession.config")

func GetConfig(ctx context.Context) *Config {
	return ctx.Value(ContextKeyConfig).(*Config)
}

var ContextKeySessionType = usersystem.ContextKey("usersession.sessiontype")

func SessionTypeContext(ctx context.Context, st usersystem.SessionType) context.Context {
	return context.WithValue(ctx, ContextKeySessionType, st)
}
func GetSessionType(ctx context.Context) usersystem.SessionType {
	return ctx.Value(ContextKeySessionType).(usersystem.SessionType)
}

var CommandActiveSessionManagerConfig = herbsystem.Command("activesessionmanagerconfig")

func WrapActiveSessionManagerConfig(h func(usersystem.SessionType) (*Config, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandActiveSessionManagerConfig
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		config, err := h(GetSessionType(ctx))
		if err != nil {
			return err
		}
		if config != nil {
			ctx = context.WithValue(ctx, ContextKeyConfig, config)
		}
		return next(ctx)
	}
	return a
}

func ExecActiveSessionManagerConfig(s *usersystem.UserSystem, st usersystem.SessionType) (*Config, error) {
	ctx := context.WithValue(s.Context, ContextKeyConfig, &Config{Supported: false})
	ctx = SessionTypeContext(ctx, st)
	ctx, err := s.System.ExecActions(ctx, CommandActiveSessionManagerConfig)
	if err != nil {
		return nil, err
	}
	return GetConfig(ctx), nil
}
