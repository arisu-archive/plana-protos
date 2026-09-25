package protos

type BannerLog struct {
	Id        int64  `json:",omitempty,omitzero"`
	ShowCount int32  `json:",omitempty,omitzero"`
	SlotOrder int32  `json:",omitempty,omitzero"`
	ClickTime MxTime `json:",omitempty,omitzero"`
}
