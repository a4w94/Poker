package apiserver

import (
	parameter "dbAPI/Parameter/28gang"
	middleware "dbAPI/middleware/28gang"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitPage_28gang(ctx *gin.Context) {
	tmp := middleware.ClientSearchThmRate()
	result, _ := json.Marshal(tmp)
	input := string(result)
	fmt.Println(input)
	ctx.HTML(http.StatusOK, "28gang.html", tmp)
}

func ThmRateAPI(ctx *gin.Context) {
	fmt.Println("GET ThmRate")

	tmp := middleware.ClientSearchThmRate()

	ctx.JSON(http.StatusOK, tmp)

}

func CalThmRTPAPI(ctx *gin.Context) {
	fmt.Println("API:計算理論ＲＴＰ")
	tmp := make(map[string]interface{})
	ctx.BindJSON(&tmp)

	var odd parameter.GameBetArea

	structType := reflect.TypeOf(odd)
	oddValue := reflect.ValueOf(&odd)
	for i := 0; i < structType.NumField(); i++ {

		tag := structType.Field(i).Name

		if tmp[tag] != nil {
			inputValue, _ := strconv.ParseFloat(tmp[tag].(string), 64)
			oddValue.Elem().FieldByName(tag).SetFloat(inputValue)

		}

		//
	}

	fmt.Println("inputodd: ", odd)
	thmrtp := middleware.ClientCalThmRTP(odd)

	ctx.JSON(http.StatusOK, thmrtp)
}

func CalRangeRateAPI(ctx *gin.Context) {
	fmt.Println("API:計算ＲＴＰ破百比例")

	tmp := make(map[string]interface{})
	ctx.BindJSON(&tmp)
	dataAmount, _ := strconv.Atoi(tmp["dataAmount"].(string))
	round, _ := strconv.Atoi(tmp["round"].(string))

	var odd parameter.GameBetArea

	structType := reflect.TypeOf(odd)
	oddValue := reflect.ValueOf(&odd)
	for i := 0; i < structType.NumField(); i++ {

		tag := structType.Field(i).Name
		if tmp[tag] != nil {
			inputValue, _ := strconv.ParseFloat(tmp[tag].(string), 64)
			oddValue.Elem().FieldByName(tag).SetFloat(inputValue)
		}
	}

	fmt.Println("inputodd: ", odd)

	rate := middleware.ClientCalOutofRangeRate(round, dataAmount, odd)

	ctx.JSON(http.StatusOK, rate)
}
