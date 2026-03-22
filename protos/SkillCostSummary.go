package protos

type SkillCostSummary struct {
	InitialCost           float32 `json:",omitempty,omitzero"`
	CostPerFrameSnapshots CostRegenSnapshotCollection
	CostAddSnapshots      []SkillCostAddSnapshot
	CostUseSnapshots      []SkillCostUseSnapshot
}
