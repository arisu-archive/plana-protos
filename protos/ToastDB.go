package protos

type ToastDB struct {
	UniqueId  int64  `json:",omitempty,omitzero"`
	Text      string `json:",omitempty,omitzero"`
	ToastId   string `json:",omitempty,omitzero"`
	BeginDate MxTime `json:",omitempty,omitzero"`
	EndDate   MxTime `json:",omitempty,omitzero"`
	LifeTime  int32  `json:",omitempty,omitzero"`
	Delay     int32  `json:",omitempty,omitzero"`
}
