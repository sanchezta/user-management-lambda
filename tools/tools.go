package tools

import (
	"fmt"
	"time"
)


func DateMySQL () string{

	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Hour(), t.Day(), t.Minute(), t.Second())
}
