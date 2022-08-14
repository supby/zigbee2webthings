package apiservice

import (
	"github.com/gin-gonic/gin"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/supby/zigbee2webthings/api/restapi/operations"
)

type WebThingsApiService struct{}

func (s *WebThingsApiService) Healthy() bool {
	return true
}

func (s *WebThingsApiService) GetThings(ctx *gin.Context) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceID(ctx *gin.Context, params *operations.GetThingsDeviceIDParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceIDActions(ctx *gin.Context, params *operations.GetThingsDeviceIDActionsParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceIDProperties(ctx *gin.Context, params *operations.GetThingsDeviceIDPropertiesParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceIDPropertiesPropertyName(ctx *gin.Context, params *operations.GetThingsDeviceIDPropertiesPropertyNameParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) PostThingsDeviceIDActions(ctx *gin.Context, params *operations.PostThingsDeviceIDActionsParams) *api.Response {
	panic("not implemented")
}
