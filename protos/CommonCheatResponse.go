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
	ScenarioHistoryIds      []int64
	ScenarioGroupHistoryDBs []*ScenarioGroupHistoryDB
	EmblemDBs               []*EmblemDB
	StudentFrameDBs         []*StudentFrameDB
	AttendanceBookRewards   []*AttendanceBookReward
	AttendanceHistoryDBs    []*AttendanceHistoryDB
	StickerDBs              []*StickerDB
	MemoryLobbyIds          []int64
	ScenarioCollectionDBs   []*ScenarioCollectionDB
	SNSPostDBs              []*SNSPostDB
	CheatFlags              CheatFlags `json:",omitempty,omitzero"`
	DebugPopupMessage       string
}
