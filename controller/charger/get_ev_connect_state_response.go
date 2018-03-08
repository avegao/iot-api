package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

const (
	evConnectStateConnectedString    evConnectState = "CONNECTED"
	evConnectStateUnknownString      evConnectState = "UNKNOWN"
	evConnectStateNotConnectedString evConnectState = "NOT_CONNECTED"
)

type evConnectState string

type getEvConnectStateResponse struct {
	State evConnectState `json:"state"`
}

func newResponseFromGetEvConnectState(grpcResponse pb.GetEvConnectStateResponse) getEvConnectStateResponse {
	return getEvConnectStateResponse{
		State: evConnectState(grpcResponse.GetState().String()),
	}
}
