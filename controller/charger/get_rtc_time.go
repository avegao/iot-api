package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getRtcTimeResponse struct {
	RtcTime string `json:"rtc_time"`
}

func newResponseFromGetRtcTime(grpcResponse pb.GetRtcTimeResponse) getRtcTimeResponse {
	return getRtcTimeResponse{
		RtcTime: grpcResponse.GetRtcTime(),
	}
}
