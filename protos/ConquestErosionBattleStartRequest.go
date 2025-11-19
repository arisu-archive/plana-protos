package protos

type ConquestErosionBattleStartRequest struct {
	RequestPacket
	EventContentId     int64             `json:",omitempty,omitzero"`
	ConquestObjectDBId int64             `json:",omitempty,omitzero"`
	UseManageEchelon   bool              `json:",omitempty,omitzero"`
	EchelonNumber      int64             `json:",omitempty,omitzero"`
	ClanAssistUseInfo  ClanAssistUseInfo `json:",omitempty,omitzero"`
}
