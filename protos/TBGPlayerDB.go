package protos

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/arisu-archive/mapx"
)

type TBGPlayerDB struct {
	Location             *TBGPlayerDBLocationHex         `json:"loc,omitempty,omitzero"`
	EventContentId       int64                           `json:"ecid,omitempty,omitzero"`
	HitPoint             int32                           `json:"hp,omitempty,omitzero"`
	DiceId               int64                           `json:"dice,omitempty,omitzero"`
	DiceProbModifyParams *mapx.OrderedMap[string, int32] `json:"diceparams"`
	Items                []*TBGItemDB                    `json:"itm"`
	TemporaryItem        *TBGItemDB                      `json:"tempitm,omitempty,omitzero"`
	ItemEffects          []*TBGItemEffectDB              `json:"eff"`
}

type TBGPlayerDBLocationHex HexLocation

func (v TBGPlayerDBLocationHex) MarshalJSON() ([]byte, error) {
	out, err := json.Marshal(fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z))
	if err != nil {
		return nil, fmt.Errorf("TBGPlayerDBLocationHex: %w", err)
	}
	return out, nil
}

func (v *TBGPlayerDBLocationHex) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TBGPlayerDBLocationHex: %w", err)
	}
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return fmt.Errorf("TBGPlayerDBLocationHex: expected 3 components, got %d in %q", len(parts), s)
	}
	dst := [3]*int32{&v.X, &v.Y, &v.Z}
	for i, p := range parts {
		n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 32)
		if err != nil {
			return fmt.Errorf("TBGPlayerDBLocationHex: parse component %d: %w", i, err)
		}
		*dst[i] = int32(n)
	}
	return nil
}
