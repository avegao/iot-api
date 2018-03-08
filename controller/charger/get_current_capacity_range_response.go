package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getCurrentCapacityRangeInAmpsResponse struct {
	MinAmps int32 `json:"min_amps"`
	MaxAmps int32 `json:"max_amps"`
}

func newResponseFromGetCurrentCapacityRangeInAmps(grpcResponse pb.GetCurrentCapacityRangeInAmpsResponse) getCurrentCapacityRangeInAmpsResponse {
	return getCurrentCapacityRangeInAmpsResponse{
		MinAmps: grpcResponse.GetMinAmps(),
		MaxAmps: grpcResponse.GetMaxAmps(),
	}
}
