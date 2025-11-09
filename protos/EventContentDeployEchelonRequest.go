package protos

type EventContentDeployEchelonRequest struct {
	RequestPacket
	EventContentId   int64      `json:",omitempty,omitzero"`
	StageUniqueId    int64      `json:",omitempty,omitzero"`
	DeployedEchelons []HexaUnit `json:",omitempty,omitzero"`
}
