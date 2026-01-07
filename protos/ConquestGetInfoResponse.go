package protos

type ConquestGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB           ConquestInfoDB          `json:",omitempty,omitzero"`
	ConquestedTileDBs        []ConquestTileDB        `json:",omitempty,omitzero"`
	ConquestObjectDBsWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	ConquestEchelonDBs       []ConquestEchelonDB     `json:",omitempty,omitzero"`
	DifficultyToStepDict     map[string]int32        `json:",omitempty,omitzero"`
	IsFirstEnter             bool                    `json:",omitempty,omitzero"`
	DisplayInfos             []ConquestDisplayInfo   `json:",omitempty,omitzero"`
}
