package main

import (
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

func sInSlice(s string, slice []string) bool {
	for _, t := range slice {
		if s == t {
			return true
		}
	}
	return false
}

var logger *logrus.Logger
var remParRegex = regexp.MustCompile(`\([^)]*\)`)
var remHexRegex = regexp.MustCompile(`\+0x[0-9a-fA-F]+`)

func cleanLogLine(s string) string {
	return strings.TrimSpace(remHexRegex.ReplaceAllString(remParRegex.ReplaceAllLiteralString(s, ""), ""))
}

func DebugStackFormatter() string {
	stackTrace := debug.Stack()
	stackLines := strings.Split(string(stackTrace), "\n")
	var stack []string
	for i := 1; i < len(stackLines); i++ {
		line := stackLines[i]
		if sInSlice(line, []string{"runtime", "testing", "logrus", "logging"}) {
			continue
		}
		fileName := ""
		funcName := ""
		if i < (len(stackLines) - 1) {
			if strings.Contains(stackLines[i+1], ".go") {
				fileName = strings.TrimSpace(stackLines[i+1])
				if strings.Contains(fileName, "/") {
					parts := strings.Split(fileName, "/")
					nLineFunc := parts[len(parts)-1]
					fileName = cleanLogLine(nLineFunc)
				}
			}
		}
		if !strings.Contains(stackLines[i], ".go") {
			funcName = cleanLogLine(stackLines[i])
			if strings.Contains(funcName, "/") {
				parts := strings.Split(funcName, "/")
				funcName = parts[len(parts)-1]
			}
			funcName = strings.TrimSpace(funcName)
		}
		if fileName != "" {
			funcName += " " + fileName
		}
		if funcName != "" && fileName != "" && !sInSlice(funcName, stack) && !sInSlice(fileName, stack) {
			stack = append([]string{funcName}, stack...)
		}
	}
	return strings.Join(stack, " ~> ")
}

func init() {
	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(_ *runtime.Frame) (function, file string) {
			return "", DebugStackFormatter()
		},
		DisableQuote:     true,
		DisableTimestamp: true,
		DisableColors:    true,
	})
}
