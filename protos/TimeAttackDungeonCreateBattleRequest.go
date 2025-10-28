package protos

type TimeAttackDungeonCreateBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
}
