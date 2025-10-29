package protos

type CraftSelectNodeResponse struct {
	ResponsePacket
	SelectedNodeDB CraftNodeDB `json:",omitempty,omitzero"`
}
