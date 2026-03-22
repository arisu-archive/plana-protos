package protos

type MiniGameCCGCompleteGameResponse struct {
	ResponsePacket
	OldSaveDB      MiniGameCCGSaveDB
	ParcelResultDB ParcelResultDB
	RewardParcels  []ParcelInfo
}
