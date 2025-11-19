package protos

type EventContentRetreatResponse struct {
	ResponsePacket
	ReleasedEchelonNumbers []int64        `json:",omitempty,omitzero"`
	ParcelResultDB         ParcelResultDB `json:",omitempty,omitzero"`
}
