package protos

type IrcServerConfig struct {
	HostAddress string `json:",omitempty,omitzero"`
	Port        int32  `json:",omitempty,omitzero"`
	Password    string `json:",omitempty,omitzero"`
}
