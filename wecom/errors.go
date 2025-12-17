package wecom

type Error struct {
	Code    int
	Message string
	Raw     []byte
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}
