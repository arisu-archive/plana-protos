package protos

type CampaignRetreatResponse struct {
	ResponsePacket
	ReleasedEchelonNumbers []int64        `json:",omitempty,omitzero"`
	ParcelResultDB         ParcelResultDB `json:",omitempty,omitzero"`
}
