package protos

type ContentSweepMultiSweepResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClearParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	BonusParcels []ParcelInfo `json:",omitempty,omitzero"`
	EventContentBonusParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDBs []CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
