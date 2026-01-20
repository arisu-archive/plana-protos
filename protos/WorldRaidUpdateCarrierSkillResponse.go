package protos

type WorldRaidUpdateCarrierSkillResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB   `json:",omitempty,omitzero"`
	CarrierSkills  map[string]int32 `json:",omitempty,omitzero"`
}
