package protos

type SchoolDungeonBattleResultRequest struct {
	RequestPacket
	StageUniqueId      int64         `json:",omitempty,omitzero"`
	PassCheckCharacter bool          `json:",omitempty,omitzero"`
	Summary            BattleSummary `json:",omitempty,omitzero"`
}
