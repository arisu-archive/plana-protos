package protos

type ContentSweepMultiSweepResponse struct {
	ResponsePacket
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             *ParcelResultDB `json:",omitempty,omitzero"`
	CampaignStageHistoryDBs  []CampaignStageHistoryDB
}
