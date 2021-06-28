package utils

import (
	"log"
	"strconv"
)

func StringToInt(num string) int {
	numInt,err:=strconv.Atoi(num)
	if err!=nil {
		log.Printf("string to int error %s",err)
	}
	log.Printf("string to int sucess")
	return numInt
}
