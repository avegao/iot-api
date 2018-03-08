package chargerController

import (
	pb "github.com/avegao/iot-api/resource/grpc/openevse"
)

type getVersionResponse struct {
	ProtocolVersion string `json:"protocol_version"`
	FirmwareVersion string `json:"firmware_version"`
}

func newResponseFromGetVersion(grpcResponse pb.GetVersionResponse) getVersionResponse {
	return getVersionResponse{
		ProtocolVersion: grpcResponse.GetProtocolVersion(),
		FirmwareVersion: grpcResponse.GetFirmwareVersion(),
	}
}
