package util

import "os"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Hostname() string {
	name, err := os.Hostname()
	Check(err)
	return name
}
