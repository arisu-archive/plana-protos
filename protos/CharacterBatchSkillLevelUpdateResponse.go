package protos

type CharacterBatchSkillLevelUpdateResponse struct {
	ResponsePacket
	CharacterDB    CharacterDB    `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
