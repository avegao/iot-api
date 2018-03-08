package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getSettingsResponse struct {
	Amperes int32    `json:"amperes"`
	Flags   []string `json:"flags"`
}

func newResponseFromGetSettings(grpcResponse pb.GetSettingsResponse) getSettingsResponse {
	return getSettingsResponse{
		Amperes: grpcResponse.GetAmperes(),
		Flags:   grpcResponse.GetFlags(),
	}
}
