package fastjson

import "errors"

func ParseString(input string) (string, error) {
	if len(input) < 2 {
		return "", errors.New("invalid JSON string: too short")
	}
	if input[0] != '"' || input[len(input)-1] != '"' {
		 return "", errors.New("invalid JSON string: missing quotes")
	}
	 return input[1 : len(input) - 1], nil
}