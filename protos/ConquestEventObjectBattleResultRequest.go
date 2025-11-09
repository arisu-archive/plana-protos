package protos

type ConquestEventObjectBattleResultRequest struct {
	RequestPacket
	EventContentId     int64         `json:",omitempty,omitzero"`
	ConquestObjectDBId int64         `json:",omitempty,omitzero"`
	BattleSummary      BattleSummary `json:",omitempty,omitzero"`
}
