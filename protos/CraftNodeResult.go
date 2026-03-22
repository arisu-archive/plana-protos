package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CraftNodeResult struct {
	NodeTier   flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	ParcelInfo ParcelInfo
}
