package protos

type ConquestMainStoryConquerResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB        `json:",omitempty,omitzero"`
	ConquestTileDB ConquestTileDB        `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB        `json:",omitempty,omitzero"`
	DisplayInfos   []ConquestDisplayInfo `json:",omitempty,omitzero"`
}
