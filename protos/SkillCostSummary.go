package protos

type SkillCostSummary struct {
	InitialCost           float32                      `json:",omitempty,omitzero"`
	CostPerFrameSnapshots *CostRegenSnapshotCollection `json:",omitempty,omitzero"`
	CostAddSnapshots      []SkillCostAddSnapshot
	CostUseSnapshots      []SkillCostUseSnapshot
}
