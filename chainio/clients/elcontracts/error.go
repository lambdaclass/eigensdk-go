package elcontracts

import "fmt"

type Error struct {
	code        int
	message     string
	description string
	cause       error
}

func (e Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s(%d) - %s: %s", e.message, e.code, e.description, e.cause.Error())
	}
	return fmt.Sprintf("%s(%d) - %s", e.message, e.code, e.description)
}

func (e Error) Unwrap() error {
	return e.cause
}

func BindingError(bindingName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", bindingName)
	return Error{
		code:        0,
		message:     "Binding error",
		description: errDescription,
		cause:       errorCause,
	}
}

func MissingContractError(contractName string) Error {
	errDescription := fmt.Sprintf("%s contract not provided", contractName)
	return Error{
		code:        1,
		message:     "Missing needed contract",
		description: errDescription,
		cause:       nil,
	}
}

func NestedError(functionName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", functionName)
	return Error{
		code:        2,
		message:     "Nested error",
		description: errDescription,
		cause:       errorCause,
	}
}

func CommonErrorMissingContract(contractName string) string {
	return fmt.Sprintf("Missing needed contract(1) - %s contract not provided", contractName)
}
