package protos

type CafeRankUpResponse struct {
	ResponsePacket
	CafeDB          CafeDB          `json:",omitempty,omitzero"`
	ParcelResultDB  ParcelResultDB  `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
