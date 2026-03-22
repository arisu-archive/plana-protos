package protos

import (
	"github.com/arisu-archive/mapx"
)

type CharacterSetFavoritesRequest struct {
	RequestPacket
	ActivateByServerIds *mapx.OrderedMap[int64, bool]
}
