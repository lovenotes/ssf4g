package gamehandler

import (
	"ssf4g/protobuf/resource-proto"

	"golang.org/x/net/context"
)

type GameIntrService struct {
}

func (gameintrservice *GameIntrService) ZoneStatusGet(ctx context.Context, reqproto *resourceproto.ZoneStatusGetReqProto) (*resourceproto.ZoneStatusGetRespProto, error) {
	return nil, nil
}
