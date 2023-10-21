package helper

type ErrorMessage struct {
	VI string `json:"vi"`
	EN string `json:"en"`
}

type Err struct {
	Raw          error        `json:"-"`
	HTTPCode     int          `json:"http_code"`
	ErrorCode    int          `json:"error_code"`
	ErrorMessage ErrorMessage `json:"error_message"`
}

func (e Err) Error() string {
	if e.Raw == nil {
		return ""
	}

	return e.Raw.Error()
}
