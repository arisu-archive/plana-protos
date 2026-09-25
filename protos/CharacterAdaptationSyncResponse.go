package protos

type CharacterAdaptationSyncResponse struct {
	ResponsePacket
	CharacterAdaptationDBs []*CharacterAdaptationDB
}
