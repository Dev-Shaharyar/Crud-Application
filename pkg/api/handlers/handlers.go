package handlers

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(uh *UserHandler) *Handlers {
	return &Handlers{
		UserHandler: uh,
	}
}
