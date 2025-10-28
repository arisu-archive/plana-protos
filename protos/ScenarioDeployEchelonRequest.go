package protos

type ScenarioDeployEchelonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	DeployedEchelons []HexaUnit `json:",omitempty,omitzero"`
}
