package protos

type CraftAutoBeginProcessResponse struct {
	ResponsePacket
	CraftInfoDBs []CraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
