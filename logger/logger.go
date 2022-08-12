package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var mod = "VERSC"

func traceM(name string, msg ...any) {
	if viper.GetBool("trace") {
		fmt.Fprintf(os.Stderr, "[ %s:TRACE ]: %s\n", name, fmt.Sprint(msg...))
	}
}

func Trace(msg ...any) {
	traceM(mod, msg...)
}

func debugM(name string, msg ...any) {
	if viper.GetBool("debug") || viper.GetBool("trace") {
		fmt.Fprintf(os.Stderr, "[ %s:DEBUG ]: %s\n", name, fmt.Sprint(msg...))
	}
}
func Debug(msg ...any) {
	debugM(mod, msg...)
}

func infoM(name string, msg ...any) {
	if viper.GetBool("debug") || viper.GetBool("trace") || viper.GetBool("verbosity") {
		fmt.Fprintf(os.Stderr, "[ %s:INFO ]: %s\n", name, fmt.Sprint(msg...))
	}
}
func Info(msg ...any) {
	infoM(mod, msg...)
}

func warnM(name string, msg ...any) {
	fmt.Fprintf(os.Stderr, "[ %s:WARN  ]: %s\n", name, fmt.Sprint(msg...))
}

func Warn(msg ...any) {
	warnM(mod, msg...)
}

func errorM(name string, msg ...any) {
	fmt.Fprintf(os.Stderr, "[ %s:ERROR ]: %s\n", name, fmt.Sprint(msg...))
}

func Error(msg ...any) {
	errorM(mod, msg...)
}

func fatalM(name string, msg ...any) {
	fmt.Fprintf(os.Stderr, "[ %s:FATAL ]: %s\n", name, fmt.Sprint(msg...))
	os.Exit(1)
}

func Fatal(msg ...any) {
	fatalM(mod, msg...)
}

func panicM(name string, msg ...any) {
	fmt.Fprintf(os.Stderr, "[ %s:PANIC ]: %s\n", name, fmt.Sprint(msg...))
	os.Exit(1)
}

func Panic(msg ...any) {
	panicM(mod, msg...)
}
