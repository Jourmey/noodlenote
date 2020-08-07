package utils

import (
	"log"
	"strconv"
)

func Int64Parse(val string) int64 {
	v1, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return int64(v1)
}
