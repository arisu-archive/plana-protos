package protos

type MiniGameTableBoardMoveResponse struct {
	ResponsePacket
	PlayerDB       TBGPlayerDB
	SaveDB         TBGBoardSaveDB
	ParcelResultDB ParcelResultDB
}
