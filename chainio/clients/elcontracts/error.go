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

func MissingContractError(contractName string) Error {
	errDescription := fmt.Sprintf("%s contract not provided", contractName)
	return Error{1, "Missing needed contract", errDescription, nil}
}

func BindingError(bindingName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", bindingName)
	return Error{
		0,
		"Binding error",
		errDescription,
		errorCause,
	}
}

func NestedError(functionName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", functionName)
	return Error{
		2,
		"Nested error",
		errDescription,
		errorCause,
	}
}

func CommonErrorMissingContract(contractName string) string {
	return fmt.Sprintf("Missing needed contract(1) - %s contract not provided", contractName)
}
