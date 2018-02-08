package gamehandler

import (
	"ssf4g/protobuf/resource-proto"

	"golang.org/x/net/context"
)

type GameIntrService struct {
}

func (gameintrservice *GameIntrService) S2RZoneStatusGet(ctx context.Context, reqproto *resourceproto.S2RZoneStatusGetReqProto) (*resourceproto.S2RZoneStatusGetRespProto, error) {
	return nil, nil
}
