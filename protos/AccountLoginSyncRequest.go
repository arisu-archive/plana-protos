package protos

type AccountLoginSyncRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SyncProtocols []Protocol `json:",omitempty,omitzero"`
	SkillCutInOption string `json:",omitempty,omitzero"`
}
