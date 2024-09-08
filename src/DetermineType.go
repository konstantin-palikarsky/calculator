package main

import "strconv"

func DetermineType(token string) interface{} {
	if val, err := strconv.Atoi(token); err == nil {
		return val
	}
	return token
}
