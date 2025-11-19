package protos

type BattleLogSourceType int32

const (
	BattleLogSourceType_None         BattleLogSourceType = 0
	BattleLogSourceType_Normal       BattleLogSourceType = 1
	BattleLogSourceType_Ex           BattleLogSourceType = 2
	BattleLogSourceType_Public       BattleLogSourceType = 3
	BattleLogSourceType_Passive      BattleLogSourceType = 4
	BattleLogSourceType_ExtraPassive BattleLogSourceType = 5
	BattleLogSourceType_Etc          BattleLogSourceType = 6
)
