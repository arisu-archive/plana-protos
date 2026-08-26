package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/mapx"
)

type TBGEncounterDB struct {
	EncounterId                     int64                                   `json:"eid,omitempty,omitzero"`
	InvokerServerId                 int64                                   `json:"oid,omitempty,omitzero"`
	ObjectType                      int32                                   `json:"type,omitempty,omitzero"`
	ShouldDecreaseItemEffectCounter bool                                    `json:"itm,omitempty,omitzero"`
	RewardUniqueIdByIndex           *TBGEncounterDBRewardUniqueIdByIndexStr `json:"rwd"`
}

type TBGEncounterDBRewardUniqueIdByIndexStr mapx.OrderedMap[int32, int64]

func (v TBGEncounterDBRewardUniqueIdByIndexStr) MarshalJSON() ([]byte, error) {
	inner, err := json.Marshal((*mapx.OrderedMap[int32, int64])(&v))
	if err != nil {
		return nil, fmt.Errorf("TBGEncounterDBRewardUniqueIdByIndexStr marshal inner: %w", err)
	}
	out, err := json.Marshal(string(inner))
	if err != nil {
		return nil, fmt.Errorf("TBGEncounterDBRewardUniqueIdByIndexStr marshal string: %w", err)
	}
	return out, nil
}

func (v *TBGEncounterDBRewardUniqueIdByIndexStr) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TBGEncounterDBRewardUniqueIdByIndexStr unmarshal string: %w", err)
	}
	if err := json.Unmarshal([]byte(s), (*mapx.OrderedMap[int32, int64])(v)); err != nil {
		return fmt.Errorf("TBGEncounterDBRewardUniqueIdByIndexStr unmarshal inner: %w", err)
	}
	return nil
}

// TBGEncounterDBPoly is a polymorphic wrapper for TBGEncounterDB. It holds one pointer per
// concrete subclass; exactly one is non-nil after Unmarshal. The
// concrete subtype is selected from the "type" wire field (FlatData.TBGObjectType).
type TBGEncounterDBPoly struct {
	Battle      *TBGBattleEncounterDB
	Facility    *TBGFacilityEncounterDB
	Random      *TBGRandomEncounterDB
	TreasureBox *TBGTreasureBoxEncounterDB
}

func (p TBGEncounterDBPoly) MarshalJSON() ([]byte, error) {
	var (
		out []byte
		err error
	)
	switch {
	case p.Battle != nil:
		out, err = json.Marshal(p.Battle)
	case p.Facility != nil:
		out, err = json.Marshal(p.Facility)
	case p.Random != nil:
		out, err = json.Marshal(p.Random)
	case p.TreasureBox != nil:
		out, err = json.Marshal(p.TreasureBox)
	default:
		return []byte("null"), nil
	}
	if err != nil {
		return nil, fmt.Errorf("TBGEncounterDBPoly: %w", err)
	}
	return out, nil
}

func (p *TBGEncounterDBPoly) UnmarshalJSON(data []byte) error {
	var probe struct {
		Type int `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return fmt.Errorf("TBGEncounterDBPoly: probe type: %w", err)
	}
	switch probe.Type { //nolint:mnd // FlatData.TBGObjectType discriminator
	case 1, 2:
		p.Battle = &TBGBattleEncounterDB{}
		if err := json.Unmarshal(data, p.Battle); err != nil {
			return fmt.Errorf("TBGEncounterDBPoly: %w", err)
		}
		return nil
	case 4:
		p.Facility = &TBGFacilityEncounterDB{}
		if err := json.Unmarshal(data, p.Facility); err != nil {
			return fmt.Errorf("TBGEncounterDBPoly: %w", err)
		}
		return nil
	case 3:
		p.Random = &TBGRandomEncounterDB{}
		if err := json.Unmarshal(data, p.Random); err != nil {
			return fmt.Errorf("TBGEncounterDBPoly: %w", err)
		}
		return nil
	case 5:
		p.TreasureBox = &TBGTreasureBoxEncounterDB{}
		if err := json.Unmarshal(data, p.TreasureBox); err != nil {
			return fmt.Errorf("TBGEncounterDBPoly: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("TBGEncounterDBPoly: unknown type %d", probe.Type)
	}
}
