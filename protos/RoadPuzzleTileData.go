package protos

type RoadPuzzleTileData struct {
	Location HexLocation
	Rotation int32 `json:",omitempty,omitzero"`
	Uid      int64 `json:",omitempty,omitzero"`
}
