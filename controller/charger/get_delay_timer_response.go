package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getDelayTimerResponse struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func newResponseFromGetDelayTimer(grpcResponse pb.GetDelayTimerResponse) getDelayTimerResponse {
	return getDelayTimerResponse{
		StartTime: grpcResponse.GetStartTime(),
		EndTime:   grpcResponse.GetEndTime(),
	}
}
