package protos

import (
	"encoding/json"
	"fmt"
)

type CharacterSkillLevelUpdateRequest struct {
	RequestPacket
	TargetCharacterDBId int64                                            `json:",omitempty,omitzero"`
	SkillSlot           CharacterSkillLevelUpdateRequestSkillSlotDefault `json:",omitzero"`
	Level               int32                                            `json:",omitempty,omitzero"`
	ReplaceInfos        []*SelectTicketReplaceInfo
}

type CharacterSkillLevelUpdateRequestSkillSlotDefault SkillSlot

func (s CharacterSkillLevelUpdateRequestSkillSlotDefault) IsZero() bool {
	return s == CharacterSkillLevelUpdateRequestSkillSlotDefault("None")
}

func (x *CharacterSkillLevelUpdateRequest) UnmarshalJSON(data []byte) error {
	type aliasCharacterSkillLevelUpdateRequest CharacterSkillLevelUpdateRequest
	v := aliasCharacterSkillLevelUpdateRequest{
		SkillSlot: CharacterSkillLevelUpdateRequestSkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal CharacterSkillLevelUpdateRequest: %w", err)
	}
	*x = CharacterSkillLevelUpdateRequest(v)
	return nil
}
