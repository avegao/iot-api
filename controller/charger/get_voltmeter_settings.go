package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getVoltmeterSettingsResponse struct {
	Calefactor int32 `json:"calefactor"`
	Offset     int32 `json:"offset"`
}

func newResponseFromGetVoltmeterSettings(grpcResponse pb.GetVoltmeterSettingsResponse) getVoltmeterSettingsResponse {
	return getVoltmeterSettingsResponse{
		Calefactor: grpcResponse.GetCalefactor(),
		Offset:     grpcResponse.GetOffset(),
	}
}
