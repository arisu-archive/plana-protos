package protos

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB       ConquestInfoDB   `json:",omitempty,omitzero"`
	ConquestedTileDBs    []ConquestTileDB `json:",omitempty,omitzero"`
	DifficultyToStepDict map[string]int32 `json:",omitempty,omitzero"`
	IsFirstEnter         bool             `json:",omitempty,omitzero"`
}
