package protos

type TimeAttackDungeonEnterBattleResponse struct {
	ResponsePacket
	AssistCharacterDB AssistCharacterDB `json:",omitempty,omitzero"`
}
