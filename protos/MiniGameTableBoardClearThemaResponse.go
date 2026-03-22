package protos

type MiniGameTableBoardClearThemaResponse struct {
	ResponsePacket
	SaveDB         TBGBoardSaveDB
	ParcelResultDB ParcelResultDB
}
