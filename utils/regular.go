package utils

import "regexp"

func IsPhone(acc string) bool {
	return regexp.MustCompile("^1[3-9]\\d{9}$").MatchString(acc)
}

func IsEmail(acc string) bool {
	return regexp.MustCompile("^([a-z0-9_.-]+)@([\\da-z.-]+).([a-z.]{2,6})$").MatchString(acc)
}
