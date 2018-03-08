package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getChargeLimitResponse struct {
	Kwh int32 `json:"kwh"`
}

func newResponseFromGetChargeLimit(grpcResponse pb.GetChargeLimitResponse) getChargeLimitResponse {
	return getChargeLimitResponse{
		Kwh: grpcResponse.GetKwh(),
	}
}
