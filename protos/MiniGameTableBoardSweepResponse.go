package protos

type MiniGameTableBoardSweepResponse struct {
	ResponsePacket
	SaveDB         TBGBoardSaveDB
	ParcelResultDB ParcelResultDB
}
