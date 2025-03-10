package utility

import "strconv"

func MustInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}
