package protos

type TimeAttackDungeonCreateBattleRequest struct {
	RequestPacket
	IsPractice bool `json:",omitempty,omitzero"`
}
