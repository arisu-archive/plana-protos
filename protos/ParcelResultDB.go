package protos

import (
	"github.com/arisu-archive/mapx"
)

type ParcelResultDB struct {
	AccountDB                       AccountDB
	AcademyLocationDBs              []AcademyLocationDB
	AccountCurrencyDB               AccountCurrencyDB
	CharacterDBs                    []CharacterDB
	WeaponDBs                       []WeaponDB
	CostumeDBs                      []CostumeDB
	TSSCharacterDBs                 []CharacterDB
	EquipmentDBs                    *mapx.OrderedMap[int64, EquipmentDB]
	RemovedEquipmentIds             []int64
	ItemDBs                         *mapx.OrderedMap[int64, ItemDB]
	RemovedItemIds                  []int64
	FurnitureDBs                    *mapx.OrderedMap[int64, FurnitureDB]
	RemovedFurnitureIds             []int64
	IdCardBackgroundDBs             *mapx.OrderedMap[int64, IdCardBackgroundDB]
	EmblemDBs                       []EmblemDB
	StickerDBs                      []StickerDB
	MemoryLobbyDBs                  []MemoryLobbyDB
	CharacterNewUniqueIds           []int64
	SecretStoneCharacterIdAndCounts *mapx.OrderedMap[int64, int32]
	DisplaySequence                 []ParcelInfo
	ParcelForMission                []ParcelInfo
	ParcelResultStepInfoList        []ParcelResultStepInfo
	BaseAccountExp                  int64 `json:",omitempty,omitzero"`
	AdditionalAccountExp            int64 `json:",omitempty,omitzero"`
	NewbieBoostAccountExp           int64 `json:",omitempty,omitzero"`
}
