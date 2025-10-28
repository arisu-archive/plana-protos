package protos

type EventContentRetreatResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ReleasedEchelonNumbers []int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
