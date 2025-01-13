package service

import (
	"fmt"
	"log"
	"os"
)

var (
	ErrorLogger = log.New(os.Stderr, "[\x1b[31mNARAKA_OPTIMIZER_TOOL_ERROR\x1b[0m] ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger  = log.New(os.Stdout, "[\x1b[32mNARAKA_OPTIMIZER_TOOL_INFO\x1b[0m] ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(os.Stdout, "[\x1b[33mNARAKA_OPTIMIZER_TOOL_DEBUG\x1b[0m] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func NewLog(format string, v ...any) string {
	if v == nil {
		return format + "\n"
	}
	return fmt.Sprintf(format+"\n", v...)
}
