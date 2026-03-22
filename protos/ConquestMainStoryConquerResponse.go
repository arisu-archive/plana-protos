package protos

type ConquestMainStoryConquerResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	ConquestTileDB ConquestTileDB
	ConquestInfoDB ConquestInfoDB
	DisplayInfos   []ConquestDisplayInfo
}
