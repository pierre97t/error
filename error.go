package error

import "fmt"

const (
	// INF Information error type
	INF uint8 = iota
	// WRN Warning error type
	WRN
	// ERR Standard error type
	ERR
	// FTL Fatal error type
	FTL
)

// Error is the custom error structure of the application it contains all required data to process an error
type Error struct {
	Code int         `json:"code,omitempty"`
	Type uint8       `json:"type,omitempty"`
	Msg  string      `json:"message,omitempty"`
	Err  error       `json:"err,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

func (e *Error) Error() string {
	if e.Body != nil {
		if byteBody, valid := e.Body.([]byte); valid {
			e.Body = fmt.Sprintf("%s", byteBody)
		} else if strBody, valid := e.Body.(string); valid {
			e.Body = strBody
		}
	}

	err := "{"
	if e.Code > 0 {
		err += fmt.Sprintf("\"code\":%v,", e.Code)
	}
	err += fmt.Sprintf("\"type\":%v,", e.Type)
	if e.Msg != "" {
		err += fmt.Sprintf("\"message\":\"%s\",", e.Msg)
	}
	if e.Err != nil {
		err += fmt.Sprintf("\"error\":\"%s\",", e.Err.Error())
	}
	if e.Body != nil {
		err += fmt.Sprintf("\"body\":\"%v\",", e.Body)
	}
	return err[0:len(err)-1] + "}"
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*Error)
	return ok
}
