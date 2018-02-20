package portalhandler

import (
	"ssf4g/protobuf/game-proto"

	"golang.org/x/net/context"
)

type PortalIntrService struct {
}

func (portalintrservice *PortalIntrService) Ping(ctx context.Context, reqproto *gameproto.PingReqProto) (*gameproto.PingRespProto, error) {
	return nil, nil
}
