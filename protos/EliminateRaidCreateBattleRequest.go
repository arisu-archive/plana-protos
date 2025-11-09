package protos

type EliminateRaidCreateBattleRequest struct {
	RequestPacket
	RaidUniqueId  int64             `json:",omitempty,omitzero"`
	IsPractice    bool              `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
