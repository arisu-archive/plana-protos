package protos

type CampaignDeployEchelonRequest struct {
	RequestPacket
	StageUniqueId    int64      `json:",omitempty,omitzero"`
	DeployedEchelons []HexaUnit `json:",omitempty,omitzero"`
}
