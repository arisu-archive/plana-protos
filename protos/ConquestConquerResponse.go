package protos

type ConquestConquerResponse struct {
	ResponsePacket
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
	ConquestTileDB               *ConquestTileDB `json:",omitempty,omitzero"`
	ConquestInfoDB               *ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []*ConquestDisplayInfo
}
