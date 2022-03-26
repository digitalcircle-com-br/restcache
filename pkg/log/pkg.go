package log

import loglib "log"

func Log(s string, p ...interface{}) {
	loglib.Printf(s, p...)
}
