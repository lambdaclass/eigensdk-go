package elcontracts

import "fmt"

type Error struct {
	code        int
	message     string
	description string
	cause       error
	// metadata    map[string]interface{}
}

func (e Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s(%d) - %s: %s", e.message, e.code, e.description, e.cause.Error())
	} else {
		return fmt.Sprintf("%s(%d) - %s", e.message, e.code, e.description)
	}
}

func CommonErrorMissingContract(contractName string) string {
	return fmt.Sprintf("Missing needed contract(1) - %s contract not provided", contractName)
}
