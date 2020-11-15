package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"io/ioutil"

	"RNG/src/utils"
)

type NumberStruct struct {
	Min    int
	Max    int
	Amount int
}

type ItemsStruct struct {
	Item []string
}

type DrawStruct struct {
	Award        []string
	Award_amount []int
	Candidate    []string
}

func errorHandler(w http.ResponseWriter, httpStatus int, errMsg string) {
	result := map[string]interface{}{
		"response": errMsg,
	}
	utils.ResponseWithJson(w, httpStatus, result)
}

func DrawNumber(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 讀取後恢復

	var numberS NumberStruct
	if err := json.NewDecoder(r.Body).Decode(&numberS); err != nil {
		errorHandler(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if numberS.Amount == 0 {
		errorHandler(w, http.StatusBadRequest, "Enter the amount")
		return
	}
	if numberS.Max == 0 {
		errorHandler(w, http.StatusBadRequest, "Enter a max number")
		return
	}
	if numberS.Min > numberS.Max {
		errorHandler(w, http.StatusBadRequest, "Min value musts be smaller than Max value")
		return
	}

	var result []int
	result = utils.RNG(numberS.Amount, numberS.Min, uint64(numberS.Max))

	response := map[string][]int{
		"result": result,
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}

func DrawItems(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 讀取後恢復

	var itemsS ItemsStruct
	if err := json.NewDecoder(r.Body).Decode(&itemsS); err != nil {
		errorHandler(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if len(itemsS.Item) == 0 {
		errorHandler(w, http.StatusBadRequest, "Enter the Items")
		return
	}

	var num int
	var max uint64 = uint64(len(itemsS.Item) - 1)
	num = utils.RNG(1, 0, max)[0]

	response := map[string]string{
		"result": itemsS.Item[num],
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}

func Draw(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 讀取後恢復

	var drawS DrawStruct
	if err := json.NewDecoder(r.Body).Decode(&drawS); err != nil {
		errorHandler(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if len(drawS.Award) == 0 {
		errorHandler(w, http.StatusBadRequest, "Enter the Items")
		return
	}
	if len(drawS.Candidate) == 0 {
		errorHandler(w, http.StatusBadRequest, "Enter the People")
		return
	}
	if utils.SumSlice(drawS.Award_amount) > len(drawS.Candidate) {
		errorHandler(w, http.StatusBadRequest, "Amount of Award is larger then Amount of People")
		return
	}

	var response []map[string]string
	for i := 0; i < len(drawS.Award); i++ {
		award_amount := drawS.Award_amount[i]

		for j := 0; j < award_amount; j++ {
			result_map := make(map[string]string)
			max := len(drawS.Candidate) - 1
			num := int(utils.RNG(1, 0, uint64(max))[0])

			result_map["award"] = drawS.Award[i]
			result_map["candidate"] = drawS.Candidate[num]

			drawS.Candidate = append(drawS.Candidate[:num], drawS.Candidate[num+1:]...)

			response = append(response, result_map)
		}
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}
