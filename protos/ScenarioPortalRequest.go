package protos

type ScenarioPortalRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
}
