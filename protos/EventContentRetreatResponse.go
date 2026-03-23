package protos

type EventContentRetreatResponse struct {
	ResponsePacket
	ReleasedEchelonNumbers []int64
	ParcelResultDB         *ParcelResultDB `json:",omitempty,omitzero"`
}
