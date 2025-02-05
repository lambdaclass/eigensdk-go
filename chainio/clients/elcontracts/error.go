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

func NestedErrorWithFunction(functionName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error happened while calling %s", functionName)
	return Error{
		code:        2,
		message:     "Nested error",
		description: errDescription,
		cause:       errorCause,
	}
}

func NestedErrorWithDescription(description string, errorCause error) Error {
	errDescription := fmt.Sprintf("Failed to %s", description)
	return Error{
		code:        2,
		message:     "Nested error",
		description: errDescription,
		cause:       errorCause,
	}
}

func OtherError(errDescription string, errorCause error) Error {
	return Error{
		code:        3,
		message:     "Other error",
		description: errDescription,
		cause:       errorCause,
	}
}

func NoSendTxOptsFailedError(errorCause error) Error {
	return Error{
		code:        3,
		message:     "Other errors",
		description: "Failed to get no send tx opts",
		cause:       errorCause,
	}
}

func TxGenerationError(bindingName string, errorCause error) Error {
	errDescription := fmt.Sprintf("Error generating tx for %s", bindingName)
	return Error{
		code:        4,
		message:     "Tx Generation",
		description: errDescription,
		cause:       errorCause,
	}
}

func SendError(errorCause error) Error {
	return Error{
		code:        5,
		message:     "Send error",
		description: "Failed to send tx with err",
		cause:       errorCause}
}

func CommonErrorMissingContract(contractName string) string {
	return fmt.Sprintf("Missing needed contract(1) - %s contract not provided", contractName)
}
