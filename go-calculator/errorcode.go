package calculator

type ErrorCode int

const (
	ErrorCodeOK         ErrorCode = 0    // OK.
	ErrorCodeDivideZero ErrorCode = 1000 // Divide by zero.
)
