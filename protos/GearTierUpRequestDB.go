package protos

type GearTierUpRequestDB struct {
	TargetServerId int64 `json:",omitempty,omitzero"`
	AfterTier      int64 `json:",omitempty,omitzero"`
	ReplaceInfos   []SelectTicketReplaceInfo
}
