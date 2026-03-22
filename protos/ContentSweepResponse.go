package protos

type ContentSweepResponse struct {
	ResponsePacket
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             ParcelResultDB
	CampaignStageHistoryDB   CampaignStageHistoryDB
}
