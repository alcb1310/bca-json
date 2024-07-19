package types

type BCAError struct {
	Code    int   `json:"code"`
	Message error `json:"message"`
}

func (e *BCAError) Error() string {
	return e.Message.Error()
}
