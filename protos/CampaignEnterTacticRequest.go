package protos

type CampaignEnterTacticRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex int64 `json:",omitempty,omitzero"`
	EnemyIndex int64 `json:",omitempty,omitzero"`
}
