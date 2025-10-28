package protos

type BattleLogCategory int32

const (
	BattleLogCategory_None BattleLogCategory = 0
	BattleLogCategory_Damage BattleLogCategory = 1
	BattleLogCategory_Heal BattleLogCategory = 2
)
