package logging

import (
	"fmt"
	"log"
)

const (
	Debug int = iota
	Info
	Warning
	Error
)

var Level = Debug

func SetLevel(l int) error {
	if l < 0 || l > Error {
		return fmt.Errorf("Log level must be between 0 and %d", Error)
	}
	Level = l
	return nil
}

var (
	RESET       = "\033[0m"
	BLACK       = "\033[30m"        /* Black */
	RED         = "\033[31m"        /* Red */
	GREEN       = "\033[32m"        /* Green */
	YELLOW      = "\033[33m"        /* Yellow */
	BLUE        = "\033[34m"        /* Blue */
	MAGENTA     = "\033[35m"        /* Magenta */
	CYAN        = "\033[36m"        /* Cyan */
	WHITE       = "\033[37m"        /* White */
	BOLDBLACK   = "\033[1m\033[30m" /* Bold Black */
	BOLDRED     = "\033[1m\033[31m" /* Bold Red */
	BOLDGREEN   = "\033[1m\033[32m" /* Bold Green */
	BOLDYELLOW  = "\033[1m\033[33m" /* Bold Yellow */
	BOLDBLUE    = "\033[1m\033[34m" /* Bold Blue */
	BOLDMAGENTA = "\033[1m\033[35m" /* Bold Magenta */
	BOLDCYAN    = "\033[1m\033[36m" /* Bold Cyan */
	BOLDWHITE   = "\033[1m\033[37m" /* Bold White */
)

func LogDebug(message string) {
	if Level <= Debug {
		log.Printf("%s%s%s\n\n", CYAN, message, RESET)
	}
}

func LogInfo(message string) {
	if Level <= Info {
		log.Printf("%s%s%s\n\n", BOLDBLUE, message, RESET)
	}
}

func LogWarning(message string) {
	if Level <= Warning {
		log.Printf("%s%s%s\n\n", YELLOW, message, RESET)
	}
}

func LogError(err error) {
	if Level <= Error {
		log.Printf("%s%v%s\n\n", RED, err, RESET)
	}
}

func LogFatalError(err error) {
	if Level <= Error {
		log.Fatalf("%s%v%s\n\n", BOLDRED, err, RESET)
	}
}
