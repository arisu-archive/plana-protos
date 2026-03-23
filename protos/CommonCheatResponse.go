package protos

type CommonCheatResponse struct {
	ResponsePacket
	Account                 *AccountDB         `json:",omitempty,omitzero"`
	AccountCurrency         *AccountCurrencyDB `json:",omitempty,omitzero"`
	CharacterDBs            []*CharacterDB
	EquipmentDBs            []*EquipmentDB
	WeaponDBs               []*WeaponDB
	GearDBs                 []*GearDB
	CostumeDBs              []*CostumeDB
	ItemDBs                 []*ItemDB
	ScenarioHistoryDBs      []*ScenarioHistoryDB
	ScenarioGroupHistoryDBs []*ScenarioGroupHistoryDB
	EmblemDBs               []*EmblemDB
	AttendanceBookRewards   []*AttendanceBookReward
	AttendanceHistoryDBs    []*AttendanceHistoryDB
	StickerDBs              []*StickerDB
	MemoryLobbyDBs          []*MemoryLobbyDB
	ScenarioCollectionDBs   []*ScenarioCollectionDB
	CheatFlags              CheatFlags `json:",omitempty,omitzero"`
	DebugPopupMessage       string
}
