package protos

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGHexaObjectDB struct {
	ServerId                 int64                          `json:"sid,omitempty,omitzero"`
	UniqueId                 int64                          `json:"oid,omitempty,omitzero"`
	EncounterId              int64                          `json:"eid,omitempty,omitzero"`
	MapType                  flatdata.TBGThemaType          `json:"map,omitempty,omitzero"`
	Location                 *TBGHexaObjectDBLocationHex    `json:"loc,omitempty,omitzero"`
	Activated                bool                           `json:"active,omitempty,omitzero"`
	HitPoint                 *int32                         `json:"hp,omitempty,omitzero"`
	BeforeStoryOption        *int32                         `json:"bso,omitempty,omitzero"`
	EncounterCostAlreadyPaid bool                           `json:"paid,omitempty,omitzero"`
	IsFakeTreasure           *bool                          `json:"isFake,omitempty,omitzero"`
	FixRewardUniqueIdByIndex *mapx.OrderedMap[int32, int64] `json:"fixReward,omitempty"`
}

type TBGHexaObjectDBLocationHex HexLocation

func (v TBGHexaObjectDBLocationHex) MarshalJSON() ([]byte, error) {
	out, err := json.Marshal(fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z))
	if err != nil {
		return nil, fmt.Errorf("TBGHexaObjectDBLocationHex: %w", err)
	}
	return out, nil
}

func (v *TBGHexaObjectDBLocationHex) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TBGHexaObjectDBLocationHex: %w", err)
	}
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return fmt.Errorf("TBGHexaObjectDBLocationHex: expected 3 components, got %d in %q", len(parts), s)
	}
	dst := [3]*int32{&v.X, &v.Y, &v.Z}
	for i, p := range parts {
		n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 32)
		if err != nil {
			return fmt.Errorf("TBGHexaObjectDBLocationHex: parse component %d: %w", i, err)
		}
		*dst[i] = int32(n)
	}
	return nil
}
