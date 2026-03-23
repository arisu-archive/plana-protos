package protos

type CafeSummonCharacterTicketUseResponse struct {
	ResponsePacket
	CafeDB         *CafeDB `json:",omitempty,omitzero"`
	CafeDBs        []*CafeDB
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
