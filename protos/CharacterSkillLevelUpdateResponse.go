package protos

type CharacterSkillLevelUpdateResponse struct {
	ResponsePacket
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
