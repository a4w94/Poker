package client

import (
	SQL "dbAPI/DataBase"
	parameter "dbAPI/Parameter/28gang"
	dbData "dbAPI/data/28gang"
	"reflect"

	"gorm.io/gorm"
)

//依照遊戲修改

var InputOdd parameter.GameBetArea

func ClientInputOdds(odds parameter.GameBetArea) {

	tmp := &InputOdd
	*tmp = odds
}

func ClientSearchAllData(showround int, dataAmount int) map[int]parameter.GameBetArea {

	//var finalround int
	db := SQL.Connect()
	Allresult := make(map[int]parameter.GameBetArea)

	for i := 1; i > 0; i++ {
		result := SearchRateByRound(db, showround*i, dataAmount)

		dataRTP := CalDataRTP(result)

		// for _, k := range dataRTP {
		// 	fmt.Println(showround*i, k)
		// 	fmt.Println()
		// }

		outrangeTimes := CalOutOfRangeRate(dataRTP)

		Allresult[showround*i] = outrangeTimes

		if showround*(i+1) > parameter.MaxRound {
			//finalround = showround * i
			break
		}
	}

	return Allresult
}

func ClientSearchThmRate(showdecimal int) parameter.GameData {
	db := SQL.Connect()
	thmrate := SearchThmRate(db)
	return thmrate
}

func SearchRateByRound(db *gorm.DB, round int, dataAmount int) []parameter.GameData {

	var condition dbData.SearchQuery
	condition.Table = parameter.Tablename

	var tmp dbData.Column
	tmp.ColumnName = "Round"
	tmp.Judge = "="
	tmp.Value = round

	var tmp2 dbData.Column
	tmp2.ColumnName = "Shoes"
	tmp2.Judge = "<="
	tmp2.Value = dataAmount

	condition.ColumnIndex = append(condition.ColumnIndex, tmp, tmp2)

	data := dbData.FindData(db, condition)

	return data
}

func SearchThmRate(db *gorm.DB) parameter.GameData {

	var condition dbData.SearchQuery
	var tmp dbData.Column
	condition.Table = parameter.Tablename
	tmp.ColumnName = "DeskNumber"
	tmp.Judge = "="
	tmp.Value = 0
	condition.ColumnIndex = append(condition.ColumnIndex, tmp)

	data := dbData.FindData(db, condition)

	return data[0]
}

func CalDataRTP(rate []parameter.GameData) []parameter.GameBetArea {

	var result []parameter.GameBetArea
	for i := 0; i < len(rate); i++ {

		result = append(result, CalRTP(rate[i].GameBetArea))

	}
	return result

}

func CalRTP(rate parameter.GameBetArea) parameter.GameBetArea {

	var result parameter.GameBetArea
	structType := reflect.TypeOf(result)
	rateValue := reflect.ValueOf(&rate)
	rtpValue := reflect.ValueOf(&result)
	oddsValue := reflect.ValueOf(&InputOdd)

	for i := 0; i < structType.NumField(); i++ {
		tag := structType.Field(i).Name
		rate := rateValue.Elem().FieldByName(tag).Interface().(float64)
		odd := oddsValue.Elem().FieldByName(tag).Interface().(float64)
		rtp := rate * odd
		rtpValue.Elem().FieldByName(tag).SetFloat(rtp)
	}
	result.WhiteBetArea = rate.White_1*InputOdd.White_1 + rate.White_2*InputOdd.White_2 + rate.White_3*InputOdd.White_3 + rate.White_4*InputOdd.White_4

	return result

}

func CalOutOfRangeRate(input []parameter.GameBetArea) parameter.GameBetArea {

	amount := len(input)
	var outofrangeTimes parameter.GameBetArea
	for i := 0; i < len(input); i++ {

		outofrangeTimes = CalOutOfRangeTime(input[i], outofrangeTimes)
	}
	var outofrangeRate parameter.GameBetArea

	structType := reflect.TypeOf(outofrangeRate)
	outofrangeRateValue := reflect.ValueOf(&outofrangeRate)
	outofrangeTimesValue := reflect.ValueOf(&outofrangeTimes)
	for i := 0; i < structType.NumField(); i++ {
		tag := structType.Field(i).Name
		times := outofrangeTimesValue.Elem().FieldByName(tag).Interface().(float64)
		rate := outofrangeRateValue.Elem().FieldByName(tag).Interface().(float64)
		rate = times / float64(amount)
		outofrangeRateValue.Elem().FieldByName(tag).SetFloat(rate)
	}

	return outofrangeRate
}

func CalOutOfRangeTime(input parameter.GameBetArea, times parameter.GameBetArea) parameter.GameBetArea {
	//判斷ＲＴＰ是否破百
	var result = times
	structType := reflect.TypeOf(result)
	inputValue := reflect.ValueOf(&input)
	resultValue := reflect.ValueOf(&result)
	for i := 0; i < structType.NumField(); i++ {
		tag := structType.Field(i).Name
		rtp := inputValue.Elem().FieldByName(tag).Interface().(float64)
		times := resultValue.Elem().FieldByName(tag).Interface().(float64)
		if rtp >= 1 {
			times++

		}
		resultValue.Elem().FieldByName(tag).SetFloat(times)
	}
	return result

}
