package protos

type EventContentMainStageSaveDB struct {
	CampaignMainStageSaveDB
	SelectedBuffDict map[int64]int64 `json:",omitempty,omitzero"`
	CurrentAppearedBuffGroupId int64 `json:",omitempty,omitzero"`
}
