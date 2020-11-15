package utils

/*
#cgo CFLAGS: -Iincludes
#cgo LDFLAGS: -Llibs -lfortuna -lstdc++ -Wl,-rpath=./
#include "fortuna.h"
*/
import "C"
import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Panic("responseWithJson error: " + err.Error())
	}
}

func RNG(amount int, min int, max uint64) []int {
	var result []int
	for i := 0; i < amount; i++ {
		// set range
		var num int = -1
		for num < min {
			num = int(C.FortunaRandom(C.uint(max + 1)))
		}
		result = append(result, num)
	}
	return result
}

func SumSlice(num_slice []int) int {
	sum := 0
	for _, val := range num_slice {
		sum += val
	}

	return sum
}
