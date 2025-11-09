package protos

type CraftInfoDB struct {
	SlotSequence      int64         `json:",omitempty,omitzero"`
	StartTime         MxTime        `json:",omitempty,omitzero"`
	EndTime           MxTime        `json:",omitempty,omitzero"`
	CraftSlotOpenDate MxTime        `json:",omitempty,omitzero"`
	Nodes             []CraftNodeDB `json:",omitempty,omitzero"`
}
