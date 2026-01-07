package protos

type OptionDB struct {
	ArenaIsAnonymous bool                              `json:",omitempty,omitzero"`
	CafeAllowCopy    OptionService_CafeAllowCopyPreset `json:",omitempty,omitzero"`
}
