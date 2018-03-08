package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getAuthLockStateResponse struct {
	Locked bool `json:"locked"`
}

func newResponseFromGetAuthLockState(grpcResponse pb.GetAuthLockStateResponse) getAuthLockStateResponse {
	return getAuthLockStateResponse{
		Locked: grpcResponse.GetLocked(),
	}
}
