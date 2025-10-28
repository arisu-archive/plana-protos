package protos

type CurrencyTransaction struct {
	ParcelBase
	Gold int64 `json:",omitempty,omitzero"`
	Gem int64 `json:",omitempty,omitzero"`
	GemBonus int64 `json:",omitempty,omitzero"`
	GemPaid int64 `json:",omitempty,omitzero"`
	ActionPoint int64 `json:",omitempty,omitzero"`
	ArenaTicket int64 `json:",omitempty,omitzero"`
	RaidTicket int64 `json:",omitempty,omitzero"`
	WeekDungeonChaserATicket int64 `json:",omitempty,omitzero"`
	WeekDungeonChaserBTicket int64 `json:",omitempty,omitzero"`
	WeekDungeonChaserCTicket int64 `json:",omitempty,omitzero"`
	WeekDungeonFindGiftTicket int64 `json:",omitempty,omitzero"`
	WeekDungeonBloodTicket int64 `json:",omitempty,omitzero"`
	AcademyTicket int64 `json:",omitempty,omitzero"`
	SchoolDungeonATicket int64 `json:",omitempty,omitzero"`
	SchoolDungeonBTicket int64 `json:",omitempty,omitzero"`
	SchoolDungeonCTicket int64 `json:",omitempty,omitzero"`
	TimeAttackDungeonTicket int64 `json:",omitempty,omitzero"`
	MasterCoin int64 `json:",omitempty,omitzero"`
	WorldRaidTicketA int64 `json:",omitempty,omitzero"`
	WorldRaidTicketB int64 `json:",omitempty,omitzero"`
	WorldRaidTicketC int64 `json:",omitempty,omitzero"`
	ChaserTotalTicket int64 `json:",omitempty,omitzero"`
	SchoolDungeonTotalTicket int64 `json:",omitempty,omitzero"`
	EliminateTicketA int64 `json:",omitempty,omitzero"`
	EliminateTicketB int64 `json:",omitempty,omitzero"`
	EliminateTicketC int64 `json:",omitempty,omitzero"`
	EliminateTicketD int64 `json:",omitempty,omitzero"`
}
