package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getTimeLimitResponse struct {
	Limit int32 `json:"limit"`
}

func newResponseFromGetTimeLimit(grpcResponse pb.GetTimeLimitResponse) getTimeLimitResponse {
	return getTimeLimitResponse{
		Limit: grpcResponse.GetLimit(),
	}
}
