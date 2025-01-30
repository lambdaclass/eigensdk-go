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

func CreateErrorForMissingContract(contractName string) Error {
	errDescription := fmt.Sprintf("%s contract not provided", contractName)
	return Error{1, "Missing needed contract", errDescription, nil}
}

func CreateForBindingError(bindingName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", bindingName)
	return Error{
		0,
		"Binding error",
		errDescription,
		errorCause,
	}
}

func CreateForNestedError(functionName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", functionName)
	return Error{
		2,
		"Nested error",
		errDescription,
		errorCause,
	}
}

func CreateNoSendTxOptsFailedError(errorCause error) Error {
	return Error{3, "Other errors", "Failed to get no send tx opts", errorCause}
}

func CreateForOtherError(errDescription string, errorCause error) Error {
	return Error{
		3,
		"Other error",
		errDescription,
		errorCause,
	}
}

func CreateForTxGenerationError(bindingName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error generating tx for %s", bindingName)
	return Error{
		4,
		"Tx Generation",
		errDescription,
		errorCause,
	}
}

func CreateForSendError(errorCause error) Error {
	return Error{5, "Send error", "Failed to send tx with err", errorCause}
}

func CommonErrorMissingContract(contractName string) string {
	return fmt.Sprintf("Missing needed contract(1) - %s contract not provided", contractName)
}
