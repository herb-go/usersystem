package usersystem

type Directive interface {
	Execute(*UserSystem) error
}

type DirectiveFunc func(*UserSystem) error

func (f DirectiveFunc) Execute(s *UserSystem) error {
	return f(s)
}

// DirectiveFactory herbsystem Directive create factory.
type DirectiveFactory func(loader func(v interface{}) error) (Directive, error)
