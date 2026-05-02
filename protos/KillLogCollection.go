package protos

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type KillLogCollection []*KillLog

func (c KillLogCollection) MarshalJSON() ([]byte, error) {
	m := make(map[string]int32, len(c))
	for _, k := range c {
		if k == nil || k.EntityId == nil {
			continue
		}
		m[strconv.Itoa(int(k.EntityId.UniqueId))] = k.Frame
	}
	out, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("KillLogCollection: %w", err)
	}
	return out, nil
}

func (c *KillLogCollection) UnmarshalJSON(data []byte) error {
	var m map[string]int32
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("KillLogCollection: %w", err)
	}
	out := make([]*KillLog, 0, len(m))
	for k, v := range m {
		id, err := strconv.Atoi(k)
		if err != nil {
			return fmt.Errorf("KillLogCollection: parse key %q: %w", k, err)
		}
		//nolint:gosec // It is safe to convert int to int32 here because the input is from a JSON map key which is a string representation of an integer.
		out = append(out, &KillLog{Frame: v, EntityId: &EntityId{UniqueId: int32(id)}})
	}
	*c = KillLogCollection(out)
	return nil
}
