package handler

type PostbackHandler struct {
}

func NewPostbackHandler() *PostbackHandler {
	return &PostbackHandler{}
}

func (h *PostbackHandler) Handle() error {
	return nil
}
