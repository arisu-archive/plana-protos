package protos

type IrcServerConfig struct {
	HostAddress string
	Port        int32 `json:",omitempty,omitzero"`
	Password    string
}
