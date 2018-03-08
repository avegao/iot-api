package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getEnergyUsageResponse struct {
	WhInSession    float32 `json:"wh_in_session"`
	KwhAccumulated float32 `json:"kwh_accumulated"`
}

func newResponseFromGetEnergyUsage(grpcResponse pb.GetEnergyUsageResponse) getEnergyUsageResponse {
	return getEnergyUsageResponse{
		WhInSession:    grpcResponse.GetWhInSession(),
		KwhAccumulated: grpcResponse.GetKwhAccumulated(),
	}
}
