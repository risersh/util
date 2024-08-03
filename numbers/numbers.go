package numbers

import (
	"log"
	"strconv"
)

func GetInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Printf("GetInt32 error parsing %s: %v", str, err)
		return 0
	}

	return n
}

func GetFloat64(str string) float64 {
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Printf("GetFloat64 error parsing %s: %v", str, err)
		return 0
	}

	return n
}
