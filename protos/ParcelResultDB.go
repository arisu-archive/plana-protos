package protos

import (
	"github.com/arisu-archive/mapx"
)

type ParcelResultDB struct {
	AccountDB                       AccountDB                                   `json:",omitempty,omitzero"`
	AcademyLocationDBs              []AcademyLocationDB                         `json:",omitempty,omitzero"`
	AccountCurrencyDB               AccountCurrencyDB                           `json:",omitempty,omitzero"`
	CharacterDBs                    []CharacterDB                               `json:",omitempty,omitzero"`
	WeaponDBs                       []WeaponDB                                  `json:",omitempty,omitzero"`
	CostumeDBs                      []CostumeDB                                 `json:",omitempty,omitzero"`
	TSSCharacterDBs                 []CharacterDB                               `json:",omitempty,omitzero"`
	EquipmentDBs                    *mapx.OrderedMap[int64, EquipmentDB]        `json:",omitempty,omitzero"`
	RemovedEquipmentIds             []int64                                     `json:",omitempty,omitzero"`
	ItemDBs                         *mapx.OrderedMap[int64, ItemDB]             `json:",omitempty,omitzero"`
	RemovedItemIds                  []int64                                     `json:",omitempty,omitzero"`
	FurnitureDBs                    *mapx.OrderedMap[int64, FurnitureDB]        `json:",omitempty,omitzero"`
	RemovedFurnitureIds             []int64                                     `json:",omitempty,omitzero"`
	IdCardBackgroundDBs             *mapx.OrderedMap[int64, IdCardBackgroundDB] `json:",omitempty,omitzero"`
	EmblemDBs                       []EmblemDB                                  `json:",omitempty,omitzero"`
	StickerDBs                      []StickerDB                                 `json:",omitempty,omitzero"`
	MemoryLobbyDBs                  []MemoryLobbyDB                             `json:",omitempty,omitzero"`
	CharacterNewUniqueIds           []int64                                     `json:",omitempty,omitzero"`
	SecretStoneCharacterIdAndCounts *mapx.OrderedMap[int64, int32]              `json:",omitempty,omitzero"`
	DisplaySequence                 []ParcelInfo                                `json:",omitempty,omitzero"`
	ParcelForMission                []ParcelInfo                                `json:",omitempty,omitzero"`
	ParcelResultStepInfoList        []ParcelResultStepInfo                      `json:",omitempty,omitzero"`
	BaseAccountExp                  int64                                       `json:",omitempty,omitzero"`
	AdditionalAccountExp            int64                                       `json:",omitempty,omitzero"`
	NewbieBoostAccountExp           int64                                       `json:",omitempty,omitzero"`
}
