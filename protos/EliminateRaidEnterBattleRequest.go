package protos

type EliminateRaidEnterBattleRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
	RaidUniqueId int64 `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	EchelonId int64 `json:",omitempty,omitzero"`
	AssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
}
