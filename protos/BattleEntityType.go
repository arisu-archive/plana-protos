package protos

type BattleEntityType int32

const (
	BattleEntityType_None BattleEntityType = 0
	BattleEntityType_Character BattleEntityType = 16777216
	BattleEntityType_SkillActor BattleEntityType = 33554432
	BattleEntityType_Obstacle BattleEntityType = 67108864
	BattleEntityType_Point BattleEntityType = 134217728
	BattleEntityType_Projectile BattleEntityType = 268435456
	BattleEntityType_EffectArea BattleEntityType = 536870912
	BattleEntityType_Supporter BattleEntityType = 1073741824
	BattleEntityType_BattleItem BattleEntityType = 2147483647
)
