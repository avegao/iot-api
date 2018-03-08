package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getOverTemperatureThresholdsResponse struct {
	Ambient float32 `json:"ambient"`
	Ir      float32 `json:"ir"`
}

func newResponseFromOverTemperatureThresholds(grpcResponse pb.GetOverTemperatureThresholdsResponse) getOverTemperatureThresholdsResponse {
	return getOverTemperatureThresholdsResponse{
		Ambient:       grpcResponse.GetAmbient(),
		Ir:   grpcResponse.GetIr(),
	}
}
