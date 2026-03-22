package protos

type ContentSweepMultiSweepResponse struct {
	ResponsePacket
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             ParcelResultDB
	CampaignStageHistoryDBs  []CampaignStageHistoryDB
}
