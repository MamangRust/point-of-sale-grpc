package errors

import (
	"encoding/json"
	"pointofsale/internal/pb"
)

func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
