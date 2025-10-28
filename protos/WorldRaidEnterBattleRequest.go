package protos

type WorldRaidEnterBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	EchelonId int64 `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	IsTicket bool `json:",omitempty,omitzero"`
	AssistUseInfos []ClanAssistUseInfo `json:",omitempty,omitzero"`
}
