package models

type ErrorResponse struct {
	Message string
	Errors  []struct {
		Path    string
		Message string
	}
}
