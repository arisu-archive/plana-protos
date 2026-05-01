package protos

import (
	"encoding/json"
	"fmt"
)

type CharacterSkillLevelUpdateRequest struct {
	RequestPacket
	TargetCharacterDBId int64                                             `json:",omitempty,omitzero"`
	SkillSlot           CharacterSkillLevelUpdateRequest_SkillSlotDefault `json:",omitzero"`
	Level               int32                                             `json:",omitempty,omitzero"`
	ReplaceInfos        []*SelectTicketReplaceInfo
}

type CharacterSkillLevelUpdateRequest_SkillSlotDefault SkillSlot

func (s CharacterSkillLevelUpdateRequest_SkillSlotDefault) IsZero() bool {
	return s == CharacterSkillLevelUpdateRequest_SkillSlotDefault("None")
}

func (x *CharacterSkillLevelUpdateRequest) UnmarshalJSON(data []byte) error {
	type aliasCharacterSkillLevelUpdateRequest CharacterSkillLevelUpdateRequest
	v := aliasCharacterSkillLevelUpdateRequest{
		SkillSlot: CharacterSkillLevelUpdateRequest_SkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal CharacterSkillLevelUpdateRequest: %w", err)
	}
	*x = CharacterSkillLevelUpdateRequest(v)
	return nil
}
