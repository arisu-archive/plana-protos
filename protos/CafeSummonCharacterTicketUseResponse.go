package protos

type CafeSummonCharacterTicketUseResponse struct {
	ResponsePacket
	CafeDB         CafeDB         `json:",omitempty,omitzero"`
	CafeDBs        []CafeDB       `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
