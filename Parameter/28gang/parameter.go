package parameter

const Tablename = "TwoEightGang_1"
const MaxRound = 79

type GameData struct {
	DataInfo    `json:"DataInfo"`
	GameBetArea `json:"GameRate"`
}

type DataInfo struct {
	GameType      int    `gorm:"type:Int" gorm:column:GameType" `
	DeskNumber    int    `gorm:"type:Int" gorm:column:DeskNumber" `
	UseDeck       int    `gorm:"type:Int" gorm:"column:UseDeck"`
	Shoes         int64  `gorm:"type:Int" gorm:"column:Shoes"`
	Round         int    `gorm:"type:Int" gorm:"column:Round"`
	RoundDateTime string `gorm:"type:DateTime" gorm:"column:RoundDateTime"`
	RoundResult   string ` gorm:"type:String"gorm:"column:RoundResult"`
	RemainingCard string `gorm:"type:String" gorm:"column:RemainingCard"`
}

type GameBetArea struct {
	Banker_and_Tie_Win        float64 `gorm:"type:Float" gorm:"column:Banker_and_Tie_Win"`
	Banker_Win                float64 `gorm:"type:Float" gorm:"column:Banker_Win"`
	Player_Win                float64 `gorm:"type:Float" gorm:"column:Player_Win"`
	Tie_Win                   float64 `gorm:"type:Float" gorm:"column:Tie_Win"`
	Banker_Pair               float64 `gorm:"type:Float" gorm:"column:Banker_Pair"`
	Player_Pair               float64 `gorm:"type:Float" gorm:"column:Player_Pair"`
	Banker_28gang             float64 `gorm:"type:Float" gorm:"column:Banker_28gang"`
	Player_28gang             float64 `gorm:"type:Float" gorm:"column:Player_28gang"`
	Banker_and_Player_28gang  float64 `gorm:"type:Float" gorm:"column:Banker_and_Player_28gang"`
	Banker_or_Player_28gang   float64 `gorm:"type:Float" gorm:"column:Banker_or_Player_28gang"`
	Banker_10point            float64 `gorm:"type:Float" gorm:"column:Banker_10point"`
	Player_10point            float64 `gorm:"type:Float" gorm:"column:Player_10point"`
	Banker_and_Player_10point float64 `gorm:"type:Float" gorm:"column:Banker_and_Player_10point"`
	Banker_or_Player_10point  float64 `gorm:"type:Float" gorm:"column:Banker_or_Player_10point"`
	Point_1                   float64 `gorm:"type:Float" gorm:"column:Point_1"`
	Point_2                   float64 `gorm:"type:Float" gorm:"column:Point_2"`
	Point_3                   float64 `gorm:"type:Float" gorm:"column:Point_3"`
	Point_4                   float64 `gorm:"type:Float"gorm:"column:Point_4"`
	Point_5                   float64 `gorm:"type:Float"gorm:"column:Point_5"`
	Point_6                   float64 `gorm:"type:Float"gorm:"column:Point_6"`
	Point_7                   float64 `gorm:"type:Float"gorm:"column:Point_7"`
	Point_8                   float64 `gorm:"type:Float"gorm:"column:Point_8"`
	Point_9                   float64 `gorm:"type:Float"gorm:"column:Point_9"`
	White_1                   float64 `gorm:"type:Float"gorm:"column:White_1"`
	White_2                   float64 `gorm:"type:Float"gorm:"column:White_2"`
	White_3                   float64 `gorm:"type:Float"gorm:"column:White_3"`
	White_4                   float64 `gorm:"type:Float"gorm:"column:White_4"`
	WhiteBetArea              float64
}
