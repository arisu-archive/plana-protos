package protos

type CraftSelectNodeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SelectedNodeDB CraftNodeDB `json:",omitempty,omitzero"`
}
