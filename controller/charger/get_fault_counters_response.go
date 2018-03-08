package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getFaultCountersResponse struct {
	Gfdi       int32 `json:"gfdi"`
	NoGround   int32 `json:"no_ground"`
	StuckRelay int32 `json:"stuck_relay"`
}

func newResponseFromGetFaultCounters(grpcResponse pb.GetFaultCountersResponse) getFaultCountersResponse {
	return getFaultCountersResponse{
		Gfdi:       grpcResponse.GetGfdi(),
		NoGround:   grpcResponse.GetNoGround(),
		StuckRelay: grpcResponse.GetStuckRelay(),
	}
}
