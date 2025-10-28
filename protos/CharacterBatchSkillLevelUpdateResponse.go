package protos

type CharacterBatchSkillLevelUpdateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
