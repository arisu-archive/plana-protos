package protos

type ContentSweepResponse struct {
	ResponsePacket
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             *ParcelResultDB         `json:",omitempty,omitzero"`
	CampaignStageHistoryDB   *CampaignStageHistoryDB `json:",omitempty,omitzero"`
}
