package protos

type CampaignRetreatResponse struct {
	ResponsePacket
	ReleasedEchelonNumbers []int64
	ParcelResultDB         *ParcelResultDB `json:",omitempty,omitzero"`
}
