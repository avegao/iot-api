package openevseService

import (
	"github.com/avegao/gocondi"
	"google.golang.org/grpc"
	"github.com/avegao/iot-api/resource/grpc/openevse"
	"time"
)

type Charger struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	Host      string     `json:"host"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateConnection() (connection *grpc.ClientConn, err error) {
	var grpcOptions []grpc.DialOption

	grpcOptions = append(grpcOptions, grpc.WithInsecure())
	address := gocondi.GetContainer().GetStringParameter("openevse_address")
	connection, err = grpc.Dial(address, grpcOptions...)

	if nil != err {
		gocondi.GetContainer().GetLogger().WithError(err).Fatalf("Fail to connect with %s", address)

		return
	}

	gocondi.GetContainer().GetLogger().Debugf("gRPC connection status with %v = %s", address, connection.GetState().String())

	return
}

func CreateClient(connection *grpc.ClientConn) iot_openevse_service.OpenevseClient {
	return iot_openevse_service.NewOpenevseClient(connection)
}

func NewChargerFromGrpcResponse(response iot_openevse_service.Charger) Charger {
	var deletedAtPointer *time.Time

	if response.DeletedAt != "" {
		deletedAt, _ := time.Parse(time.RFC3339, response.DeletedAt)
		deletedAtPointer = &deletedAt
	}

	createdAt, _ := time.Parse(time.RFC3339, response.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, response.UpdatedAt)

	return Charger{
		ID:        response.Id,
		Name:      response.Name,
		Host:      response.Host,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAtPointer,
	}
}
