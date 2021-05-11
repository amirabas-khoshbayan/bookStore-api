package errors

import "net/http"

type (
	//code error
	kind uint
	serverError struct {
		kind kind
		message string
	}
)

const (
	_ kind = iota
	KindInvalid
	KindNotFound
	KindUnauthorizad
	KindUnexpected
	KindNotAllowed
)

var (
	httpErrors = map[kind]int{
		KindInvalid : http.StatusBadRequest,
		KindNotFound: http.StatusNotFound,
		KindUnauthorizad : http.StatusUnauthorized,
		KindUnexpected : http.StatusInternalServerError,
		KindNotAllowed : http.StatusMethodNotAllowed,
	}
)

//return type global error
func New(kind kind,msg string) error {
   return serverError{
   	kind : kind ,
   	message: msg,
   }
}
//return string error
func (e serverError) Error() string {
  return e.message
}
//Convert serverError to httpError and received code
func HttpError(err error) (string , int)  {
	serverErr,ok := err.(serverError)
	if !ok {
     return err.Error(),http.StatusBadRequest
	}
	code,ok := httpErrors[serverErr.kind]

	if !ok {
		return serverErr.message,http.StatusBadRequest
	}

	return serverErr.message,code
}

