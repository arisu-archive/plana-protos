package protos

type CommonCheatResponse struct {
	ResponsePacket
	Account                 AccountDB
	AccountCurrency         AccountCurrencyDB
	CharacterDBs            []CharacterDB
	EquipmentDBs            []EquipmentDB
	WeaponDBs               []WeaponDB
	GearDBs                 []GearDB
	CostumeDBs              []CostumeDB
	ItemDBs                 []ItemDB
	ScenarioHistoryDBs      []ScenarioHistoryDB
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB
	EmblemDBs               []EmblemDB
	AttendanceBookRewards   []AttendanceBookReward
	AttendanceHistoryDBs    []AttendanceHistoryDB
	StickerDBs              []StickerDB
	MemoryLobbyDBs          []MemoryLobbyDB
	ScenarioCollectionDBs   []ScenarioCollectionDB
	CheatFlags              CheatFlags `json:",omitempty,omitzero"`
	DebugPopupMessage       string
}
