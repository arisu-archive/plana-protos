package protos

type ProductStepupPackageDB struct {
	Id          int64 `json:",omitempty,omitzero"`
	CurrentStep int32 `json:",omitempty,omitzero"`
}
