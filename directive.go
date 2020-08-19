package usersystem

type Directive interface {
	Execute(*UserSystem) error
}

// DirectiveFactory herbsystem Directive create factory.
type DirectiveFactory func(loader func(v interface{}) error) (Directive, error)
