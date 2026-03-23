package protos

type PermanentRaidEnterBattleRequest struct {
	RequestPacket
	StageId       int64              `json:",omitempty,omitzero"`
	AssistUseInfo *ClanAssistUseInfo `json:",omitempty,omitzero"`
}
