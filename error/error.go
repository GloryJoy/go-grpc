package main

const (
	customErrorMessage1 = "CUSTOME_ERROR"
)

type CustomError struct {
	ErrorCode   string
	ErrorString string
}

func NewCustomError() *CustomError {
	return &CustomError{
		ErrorCode:   customErrorMessage1,
		ErrorString: "this is a custom error.",
	}
}

func (c *CustomError) GetErrorCode() string {
	return c.ErrorCode
}

func (c *CustomError) Error() string {
	return c.ErrorCode + ":" + c.ErrorString
}
