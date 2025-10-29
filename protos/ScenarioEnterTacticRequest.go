package protos

type ScenarioEnterTacticRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex int64 `json:",omitempty,omitzero"`
	EnemyIndex int64 `json:",omitempty,omitzero"`
}
