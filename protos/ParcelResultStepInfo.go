package protos

type ParcelResultStepInfo struct {
	ParcelProcessActionType ParcelProcessActionType `json:",omitempty,omitzero"`
	StepParcelDetails       []*ParcelDetail
}
