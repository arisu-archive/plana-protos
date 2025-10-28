package protos

type BattleEndType int32

const (
	BattleEndType_None BattleEndType = 0
	BattleEndType_AllNearlyDead BattleEndType = 1
	BattleEndType_TimeOut BattleEndType = 2
	BattleEndType_EscortFailed BattleEndType = 3
	BattleEndType_Clear BattleEndType = 4
)
