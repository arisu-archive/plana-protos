package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/mapx"
)

type HexaUnit struct {
	EntityId                                  int64 `json:",omitempty,omitzero"`
	HpInfos                                   *mapx.OrderedMap[int64, int64]
	DyingInfos                                *mapx.OrderedMap[int64, int64]
	BuffInfos                                 *mapx.OrderedMap[int64, int32]
	ActionCountMax                            int32 `json:",omitempty,omitzero"`
	ActionCount                               int32 `json:",omitempty,omitzero"`
	Mobility                                  int32 `json:",omitempty,omitzero"`
	StrategySightRange                        int32 `json:",omitempty,omitzero"`
	Id                                        int64 `json:",omitempty,omitzero"`
	Rotate                                    HexaUnitRotateVec
	Location                                  *HexLocation `json:",omitempty,omitzero"`
	AIDestination                             *HexLocation `json:",omitempty,omitzero"`
	IsActionComplete                          bool         `json:",omitempty,omitzero"`
	IsPlayer                                  bool         `json:",omitempty,omitzero"`
	IsFixedEchelon                            bool         `json:",omitempty,omitzero"`
	MovementOrder                             int32        `json:",omitempty,omitzero"`
	RewardParcelInfosWithDropTacticEntityType *mapx.OrderedMap[string, []*ParcelInfo]
	SkillCardHand                             *SkillCardHand `json:",omitempty,omitzero"`
	PlayAnimation                             bool           `json:",omitempty,omitzero"`
}

type HexaUnitRotateVec Vector3

func (v HexaUnitRotateVec) MarshalJSON() ([]byte, error) {
	out, err := json.Marshal([3]float32{v.X, v.Y, v.Z})
	if err != nil {
		return nil, fmt.Errorf("HexaUnitRotateVec: %w", err)
	}
	return out, nil
}

func (v *HexaUnitRotateVec) UnmarshalJSON(data []byte) error {
	var a [3]float32
	if err := json.Unmarshal(data, &a); err != nil {
		return fmt.Errorf("HexaUnitRotateVec: %w", err)
	}
	v.X = a[0]
	v.Y = a[1]
	v.Z = a[2]
	return nil
}
