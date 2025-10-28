package protos

type EliminateRaidCreateBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidUniqueId int64 `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
