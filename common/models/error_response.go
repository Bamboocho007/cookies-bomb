package models

type ErrorResponse struct {
	Message string
	Errors  []ErrorResponseItem
}

type ErrorResponseItem struct {
	Path    string
	Message string
}

func (err *ErrorResponse) Error() string {
	return err.Message
}
