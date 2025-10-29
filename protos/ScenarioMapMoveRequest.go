package protos

type ScenarioMapMoveRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
	DestPosition HexLocation `json:",omitempty,omitzero"`
}
