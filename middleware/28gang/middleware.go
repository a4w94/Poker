package middleware

import (
	SQL "dbAPI/DataBase"
	parameter "dbAPI/Parameter/28gang"
	client "dbAPI/client/28gang"
	"fmt"
)

func ClientSearch(showdecimal int) {

}

func ClientSearchThmRate() parameter.GameBetArea {
	db := SQL.Connect()
	thmrate := client.SearchThmRate(db)

	return thmrate.GameBetArea
}

func ClientCalThmRTP(odd parameter.GameBetArea) parameter.GameBetArea {

	fmt.Println("Get Thm RTP")

	client.ClientInputOdds(odd)

	thmrate := ClientSearchThmRate()

	thmrtp := client.CalRTP(thmrate)

	fmt.Println("理論ＲＴＰ")
	PrintData(thmrtp)

	return thmrtp
}

func ClientCalOutofRangeRate(showroud int, dataAmount int, odd parameter.GameBetArea) map[int]parameter.GameBetArea {

	fmt.Println("Get Out Of Range Rate")
	client.ClientInputOdds(odd)
	rangerate := client.ClientSearchAllData(showroud, dataAmount)
	return rangerate

}

//可修改
func PrintData(input parameter.GameBetArea) {
	fmt.Println("和算莊贏", input.Banker_and_Tie_Win)
	fmt.Println("莊贏", input.Banker_Win, input.Banker_Win)
	fmt.Println("閒贏", input.Player_Win, input.Player_Win)
	fmt.Println("和贏", input.Tie_Win, input.Tie_Win)

	fmt.Println("莊對", input.Banker_Pair, input.Banker_Pair)
	fmt.Println("閒對", input.Player_Pair, input.Player_Pair)

	fmt.Println("莊28槓", input.Banker_28gang)
	fmt.Println("閒28槓", input.Player_28gang)

	fmt.Println("莊鱉十", input.Banker_10point)
	fmt.Println("閒鱉十", input.Player_10point)

	fmt.Println("白板注區", input.WhiteBetArea)

}
