package protos

type AccountLoginSyncRequest struct {
	RequestPacket
	SyncProtocols    []Protocol
	SkillCutInOption string
}
