package npmdl

// Error response.
type Error struct {
	Status     string
	StatusCode int
	Message    string `json:"error"`
}

// Error string.
func (e *Error) Error() string {
	return e.Message
}
