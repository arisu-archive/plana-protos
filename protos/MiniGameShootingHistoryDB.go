package protos

type MiniGameShootingHistoryDB struct {
	EventContentId int64  `json:",omitempty,omitzero"`
	UniqueId       int64  `json:",omitempty,omitzero"`
	ArriveSection  int64  `json:",omitempty,omitzero"`
	LastUpdateDate MxTime `json:",omitempty,omitzero"`
	IsClearToday   bool   `json:",omitempty,omitzero"`
}
