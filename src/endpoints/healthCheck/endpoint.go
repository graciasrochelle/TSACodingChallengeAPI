package healthCheck

import (
	"TSACodingChallengeAPI/src/common"
	"net/http"
)

func NewEndpoint() *common.Endpoint {
	return common.NewEndpoint(Handle, nil, Encode)
}

func Handle(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error) {
	response = HealthCheckResponse{
		Status: "Active",
	}
	return response, http.StatusOK, nil
}

func Encode(w http.ResponseWriter, httpStatus int, response common.ResponseType) error {
	common.EncodeJsonResponse(w, httpStatus, response)
	return nil
}
