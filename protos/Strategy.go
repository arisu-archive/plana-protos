package protos

import (
	"encoding/json"
	"fmt"
)

type Strategy struct {
	EntityId      int64 `json:",omitempty,omitzero"`
	Rotate        StrategyRotateVec
	Id            int64        `json:",omitempty,omitzero"`
	Location      *HexLocation `json:",omitempty,omitzero"`
	PlayAnimation bool         `json:",omitempty,omitzero"`
	Activated     bool         `json:",omitempty,omitzero"`
	Values        []int32
	Index         int32 `json:",omitempty,omitzero"`
}

type StrategyRotateVec Vector3

func (v StrategyRotateVec) MarshalJSON() ([]byte, error) {
	out, err := json.Marshal([3]float32{v.X, v.Y, v.Z})
	if err != nil {
		return nil, fmt.Errorf("StrategyRotateVec: %w", err)
	}
	return out, nil
}

func (v *StrategyRotateVec) UnmarshalJSON(data []byte) error {
	var a [3]float32
	if err := json.Unmarshal(data, &a); err != nil {
		return fmt.Errorf("StrategyRotateVec: %w", err)
	}
	v.X = a[0]
	v.Y = a[1]
	v.Z = a[2]
	return nil
}
