package protos

type AccountLoginSyncRequest struct {
	RequestPacket
	SyncProtocols    []Protocol `json:",omitempty,omitzero"`
	SkillCutInOption string     `json:",omitempty,omitzero"`
}
