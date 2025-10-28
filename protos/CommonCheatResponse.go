package protos

type CommonCheatResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Account AccountDB `json:",omitempty,omitzero"`
	AccountCurrency AccountCurrencyDB `json:",omitempty,omitzero"`
	CharacterDBs []CharacterDB `json:",omitempty,omitzero"`
	EquipmentDBs []EquipmentDB `json:",omitempty,omitzero"`
	WeaponDBs []WeaponDB `json:",omitempty,omitzero"`
	GearDBs []GearDB `json:",omitempty,omitzero"`
	CostumeDBs []CostumeDB `json:",omitempty,omitzero"`
	ItemDBs []ItemDB `json:",omitempty,omitzero"`
	ScenarioHistoryDBs []ScenarioHistoryDB `json:",omitempty,omitzero"`
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
	EmblemDBs []EmblemDB `json:",omitempty,omitzero"`
	AttendanceBookRewards []AttendanceBookReward `json:",omitempty,omitzero"`
	AttendanceHistoryDBs []AttendanceHistoryDB `json:",omitempty,omitzero"`
	StickerDBs []StickerDB `json:",omitempty,omitzero"`
	MemoryLobbyDBs []MemoryLobbyDB `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []ScenarioCollectionDB `json:",omitempty,omitzero"`
	CheatFlags CheatFlags `json:",omitempty,omitzero"`
}
