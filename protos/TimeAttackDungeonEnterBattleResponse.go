package protos

type TimeAttackDungeonEnterBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AssistCharacterDB AssistCharacterDB `json:",omitempty,omitzero"`
}
