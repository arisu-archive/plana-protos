package protos

type ConquestReceiveRewardsResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestTileDBs []ConquestTileDB `json:",omitempty,omitzero"`
}
