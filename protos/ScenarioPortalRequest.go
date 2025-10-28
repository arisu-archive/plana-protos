package protos

type ScenarioPortalRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
}
