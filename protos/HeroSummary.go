package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type HeroSummary struct {
	ServerId                  int64                           `json:",omitempty,omitzero"`
	OwnerAccountId            int64                           `json:",omitempty,omitzero"`
	BattleEntityId            EntityId                        `json:",omitempty,omitzero"`
	CharacterId               int64                           `json:",omitempty,omitzero"`
	CostumeId                 int64                           `json:",omitempty,omitzero"`
	Grade                     int32                           `json:",omitempty,omitzero"`
	Level                     int32                           `json:",omitempty,omitzero"`
	PotentialStatLevel        *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	ExSkillLevel              int32                           `json:",omitempty,omitzero"`
	PublicSkillLevel          int32                           `json:",omitempty,omitzero"`
	PassiveSkillLevel         int32                           `json:",omitempty,omitzero"`
	ExtraPassiveSkillLevel    int32                           `json:",omitempty,omitzero"`
	FavorRank                 int32                           `json:",omitempty,omitzero"`
	StatSnapshotCollection    StatSnapshotCollection          `json:",omitempty,omitzero"`
	HPRateBefore              int64                           `json:",omitempty,omitzero"`
	HPRateAfter               int64                           `json:",omitempty,omitzero"`
	CrowdControlCount         int32                           `json:",omitempty,omitzero"`
	CrowdControlDuration      int32                           `json:",omitempty,omitzero"`
	EvadeCount                int32                           `json:",omitempty,omitzero"`
	DamageImmuneCount         int32                           `json:",omitempty,omitzero"`
	CrowdControlImmuneCount   int32                           `json:",omitempty,omitzero"`
	MaxAttackPower            int64                           `json:",omitempty,omitzero"`
	AverageCriticalRate       int32                           `json:",omitempty,omitzero"`
	AverageStabilityRate      int32                           `json:",omitempty,omitzero"`
	AverageAccuracyRate       int32                           `json:",omitempty,omitzero"`
	DeadFrame                 int32                           `json:",omitempty,omitzero"`
	DamageGivenAbsorbedSum    int64                           `json:",omitempty,omitzero"`
	TacticEntityType          flatdata.TacticEntityType       `json:",omitempty,omitzero"`
	GivenNumericLogs          []BattleNumericLog              `json:",omitempty,omitzero"`
	TakenNumericLogs          []BattleNumericLog              `json:",omitempty,omitzero"`
	ObstacleBattleNumericLogs []BattleNumericLog              `json:",omitempty,omitzero"`
	Equipments                []EquipmentSetting              `json:",omitempty,omitzero"`
	CharacterWeapon           *WeaponSetting                  `json:",omitempty,omitzero"`
	CharacterGear             *GearSetting                    `json:",omitempty,omitzero"`
	SkillCount                *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	KillLog                   KillLogCollection               `json:",omitempty,omitzero"`
}
