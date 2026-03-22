package protos

type BattleEntityType string

const (
	BattleEntityType_None       BattleEntityType = "None"
	BattleEntityType_Character  BattleEntityType = "Character"
	BattleEntityType_SkillActor BattleEntityType = "SkillActor"
	BattleEntityType_Obstacle   BattleEntityType = "Obstacle"
	BattleEntityType_Point      BattleEntityType = "Point"
	BattleEntityType_Projectile BattleEntityType = "Projectile"
	BattleEntityType_EffectArea BattleEntityType = "EffectArea"
	BattleEntityType_Supporter  BattleEntityType = "Supporter"
	BattleEntityType_BattleItem BattleEntityType = "BattleItem"
)
