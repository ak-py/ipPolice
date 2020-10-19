package core

import "fmt"

// Event helps standardize logs
type Event struct {
	id      int
	message string
}

const (
	//LEVEL_DEBUG = "DEBUG"
	LEVEL_INFO  = "INFO"
	LEVEL_WARN  = "WARNING"
	LEVEL_ERROR = "ERROR"
	LOG_SEP     = ": "
)

// Declare variables to store log messages as new Events
var (
	stdErrorMessage         = Event{0, LEVEL_ERROR + LOG_SEP + "%s" + LOG_SEP + "%v"}
	stdErrorMessageWithCode = Event{1, LEVEL_ERROR + LOG_SEP + "%s" + LOG_SEP + "%s: %v"}
	stdWarningMessage       = Event{2, LEVEL_WARN + LOG_SEP + "%s" + LOG_SEP + "%v"}
	stdInfoMessage          = Event{3, LEVEL_INFO + LOG_SEP + "%s" + LOG_SEP + "%v"}
	variableInfoMessage     = Event{4, LEVEL_INFO + LOG_SEP + "%s" + LOG_SEP + "%s = %#v"}
)

// Standard error log message
func LogErrorMessage(err interface{}) string {
	return fmt.Sprintf(stdErrorMessage.message, callerFuncName(2), err)
}

// Standard error log message with message code
func LogErrorMessageWithCode(errMessageCode string) string {
	errMessage := GetMessage(errMessageCode)
	return fmt.Sprintf(stdErrorMessageWithCode.message, callerFuncName(2), errMessageCode, errMessage)
}

// LogInfoMessage is a standard info message
func LogWarningMessage(argumentValue interface{}) string {
	return fmt.Sprintf(stdWarningMessage.message, callerFuncName(2), argumentValue)
}

// LogInfoMessage is a standard info message
func LogInfoMessage(argumentValue interface{}) string {
	return fmt.Sprintf(stdInfoMessage.message, callerFuncName(2), argumentValue)
}

func LogInfoVariable(argumentName string, argumentValue interface{}) string {
	return fmt.Sprintf(variableInfoMessage.message, callerFuncName(2), argumentName, argumentValue)
}
