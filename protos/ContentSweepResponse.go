package protos

type ContentSweepResponse struct {
	ResponsePacket
	ClearParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	BonusParcels []ParcelInfo `json:",omitempty,omitzero"`
	EventContentBonusParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDB CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
