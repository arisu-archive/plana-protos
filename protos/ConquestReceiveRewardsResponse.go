package protos

type ConquestReceiveRewardsResponse struct {
	ResponsePacket
	ParcelResultDB  ParcelResultDB   `json:",omitempty,omitzero"`
	ConquestInfoDB  ConquestInfoDB   `json:",omitempty,omitzero"`
	ConquestTileDBs []ConquestTileDB `json:",omitempty,omitzero"`
}
