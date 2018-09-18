package constant

type ServiceError struct {
  Inner     error  // stores the error returned by external dependencies, i.e.: KeyFunc
  ErrorCode int    // error code
  text      string // errors that do not have a valid error just have text
}

// Validation error is an error type
func (e ServiceError) Error() string {
  if e.Inner != nil {
    return e.Inner.Error()
  } else if e.text != "" {
    return e.text
  } else {
    str, ok := errStrings[e.ErrorCode]
    if true == ok {
      return str
    } else {
      return ""
    }
  }
}

const (
  Success = 0

  ErrCodeAuthorizationFailed = 1001
  ErrCodeDuplicatedLogin     = 1002

  ErrCodeInternal     = 10000
  ErrCodeInvalidParam = 10001
  ErrCodeInternalDB   = 10002

  ErrCodeNotFound   = 10111
  ErrCodeExisted    = 10112
  ErrCodeNotChanged = 10113
  ErrCodeNotMatch   = 10114

  ErrNotChanged = 10102
  ErrDuplicated = 10103

  // user validation
  ErrCodeLoginInvalidParam = 10201
  ErrCodeLoginVerifyFailed = 10201

  //error
  MissingReqParameters = 20001
  OperationFailed      = 20002
)

var (
  ErrInternal        = NewError(ErrCodeInternal)
  ErrInvalidParam    = NewError(ErrCodeInvalidParam)
  ErrNotFound        = NewError(ErrCodeNotFound)
  ErrInvalidToken    = NewError(ErrCodeAuthorizationFailed)
  ErrDuplicatedLogin = NewError(ErrCodeDuplicatedLogin)
)

func NewError(code int) *ServiceError {
  return &ServiceError{ErrorCode: code}
}

var errStrings = map[int]string{
  ErrCodeInternal:     "internal error",
  ErrCodeInvalidParam: "invalid parameters",
  ErrCodeInternalDB:   "internal error",

  ErrCodeNotFound:   "not found",
  ErrCodeExisted:    "already existed",
  ErrCodeNotChanged: "not changed",
  ErrDuplicated:     "duplicated",
  ErrNotChanged:     "not changed",


  ErrCodeAuthorizationFailed: "authorization failed",
  ErrCodeDuplicatedLogin:     "duplicated login",
}

func ErrMsgFromCode(code int) (msg string) {
  msg, ok := errStrings[code]
  if false == ok {
    msg = "unknown error code"
  }
  return
}
