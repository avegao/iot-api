package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getAmmeterSettingsResponse struct {
	CurrentScaleFactor int32 `json:"current_scale_factor"`
	CurrentOffset      int32 `json:"current_offset"`
}

func newResponseFromGetAmmeterSettings(grpcResponse pb.GetAmmeterSettingsResponse) getAmmeterSettingsResponse {
	return getAmmeterSettingsResponse{
		CurrentOffset:      grpcResponse.GetCurrentOffset(),
		CurrentScaleFactor: grpcResponse.GetCurrentScaleFactor(),
	}
}
