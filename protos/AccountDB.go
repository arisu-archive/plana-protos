package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AccountDB struct {
	ServerId                   int64                 `json:",omitempty,omitzero"`
	Nickname                   string                `json:",omitempty,omitzero"`
	CallName                   string                `json:",omitempty,omitzero"`
	DevId                      string                `json:",omitempty,omitzero"`
	State                      flatdata.AccountState `json:",omitempty,omitzero"`
	Level                      int32                 `json:",omitempty,omitzero"`
	Exp                        int64                 `json:",omitempty,omitzero"`
	Comment                    string                `json:",omitempty,omitzero"`
	LobbyMode                  int32                 `json:",omitempty,omitzero"`
	RepresentCharacterServerId int64                 `json:",omitempty,omitzero"`
	MemoryLobbyUniqueId        int64                 `json:",omitempty,omitzero"`
	LastConnectTime            MxTime                `json:",omitempty,omitzero"`
	BirthDay                   MxTime                `json:",omitempty,omitzero"`
	CallNameUpdateTime         MxTime                `json:",omitempty,omitzero"`
	PublisherAccountId         int64                 `json:",omitempty,omitzero"`
	RetentionDays              *int32                `json:",omitempty,omitzero"`
	VIPLevel                   *int32                `json:",omitempty,omitzero"`
	CreateDate                 MxTime                `json:",omitempty,omitzero"`
	UnReadMailCount            *int32                `json:",omitempty,omitzero"`
	LinkRewardDate             *MxTime               `json:",omitempty,omitzero"`
	LastReturningDate          *MxTime               `json:",omitempty,omitzero"`
}
