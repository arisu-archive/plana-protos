package protos

type CraftAutoBeginProcessResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfoDBs []CraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
