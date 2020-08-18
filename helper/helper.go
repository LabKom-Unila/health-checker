package helper

import "fmt"

func SuccessOutput() {
	fmt.Println(("\033[32m"), "Ping sucessful \u2713")
}

func ErrorOutput(err error) {
	fmt.Println(("\033[31m"), err, "\u02DF")
}
