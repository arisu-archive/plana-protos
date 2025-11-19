package protos

type GearDB struct {
	ParcelBase
	ServerId               int64 `json:",omitempty,omitzero"`
	UniqueId               int64 `json:",omitempty,omitzero"`
	Level                  int32 `json:",omitempty,omitzero"`
	Exp                    int64 `json:",omitempty,omitzero"`
	Tier                   int32 `json:",omitempty,omitzero"`
	SlotIndex              int64 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
