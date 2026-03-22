package protos

type MiniGameTableBoardMoveResponse struct {
	ResponsePacket
	PlayerDB       TBGPlayerDB
	SaveDB         TBGBoardSaveDB
	EncounterDB    TBGEncounterDB
	ParcelResultDB ParcelResultDB
}
