package protos

type ExcessiveTouch struct {
	ExcessiveTouchFrameList  []int32 `json:",omitempty,omitzero"`
	ExcessiveTouchCount      int32   `json:",omitempty,omitzero"`
	TotalExcessiveTouchFound int32   `json:",omitempty,omitzero"`
}
