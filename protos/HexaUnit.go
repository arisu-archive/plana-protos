package protos

type HexaUnit struct {
	EntityId                                  int64                   `json:",omitempty,omitzero"`
	HpInfos                                   map[int64]int64         `json:",omitempty,omitzero"`
	DyingInfos                                map[int64]int64         `json:",omitempty,omitzero"`
	BuffInfos                                 map[int64]int32         `json:",omitempty,omitzero"`
	ActionCountMax                            int32                   `json:",omitempty,omitzero"`
	ActionCount                               int32                   `json:",omitempty,omitzero"`
	Mobility                                  int32                   `json:",omitempty,omitzero"`
	StrategySightRange                        int32                   `json:",omitempty,omitzero"`
	Id                                        int64                   `json:",omitempty,omitzero"`
	Rotate                                    Vector3                 `json:",omitempty,omitzero"`
	Location                                  HexLocation             `json:",omitempty,omitzero"`
	AIDestination                             HexLocation             `json:",omitempty,omitzero"`
	IsActionComplete                          bool                    `json:",omitempty,omitzero"`
	IsPlayer                                  bool                    `json:",omitempty,omitzero"`
	IsFixedEchelon                            bool                    `json:",omitempty,omitzero"`
	MovementOrder                             int32                   `json:",omitempty,omitzero"`
	RewardParcelInfosWithDropTacticEntityType map[string][]ParcelInfo `json:",omitempty,omitzero"`
	SkillCardHand                             SkillCardHand           `json:",omitempty,omitzero"`
	PlayAnimation                             bool                    `json:",omitempty,omitzero"`
}
