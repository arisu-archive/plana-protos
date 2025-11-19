package protos

type WorldRaidBattleResultRequest struct {
	RequestPacket
	SeasonId       int64               `json:",omitempty,omitzero"`
	GroupId        int64               `json:",omitempty,omitzero"`
	UniqueId       int64               `json:",omitempty,omitzero"`
	EchelonId      int64               `json:",omitempty,omitzero"`
	IsPractice     bool                `json:",omitempty,omitzero"`
	IsTicket       bool                `json:",omitempty,omitzero"`
	Summary        BattleSummary       `json:",omitempty,omitzero"`
	AssistUseInfos []ClanAssistUseInfo `json:",omitempty,omitzero"`
}
