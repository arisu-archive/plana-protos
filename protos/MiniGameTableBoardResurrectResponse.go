package protos

type MiniGameTableBoardResurrectResponse struct {
	ResponsePacket
	PlayerDB       TBGPlayerDB
	ParcelResultDB ParcelResultDB
}
