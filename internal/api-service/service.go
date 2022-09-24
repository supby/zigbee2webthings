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

func (s *WebThingsApiService) GetThingsDeviceName(ctx *gin.Context, params *operations.GetThingsDeviceNameParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNameActions(ctx *gin.Context, params *operations.GetThingsDeviceNameActionsParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNameProperties(ctx *gin.Context, params *operations.GetThingsDeviceNamePropertiesParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNamePropertiesPropertyName(ctx *gin.Context, params *operations.GetThingsDeviceNamePropertiesPropertyNameParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) PostThingsDeviceNameActions(ctx *gin.Context, params *operations.PostThingsDeviceNameActionsParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) PutThingsDeviceNameActionsActionNameActionID(ctx *gin.Context, params *operations.PutThingsDeviceNameActionsActionNameActionIDParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) DeleteThingsDeviceNameActionsActionNameActionID(ctx *gin.Context, params *operations.DeleteThingsDeviceNameActionsActionNameActionIDParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNameActionsActionNameActionID(ctx *gin.Context, params *operations.GetThingsDeviceNameActionsActionNameActionIDParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) PutThingsDeviceNamePropertiesPropertyName(ctx *gin.Context, params *operations.PutThingsDeviceNamePropertiesPropertyNameParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNameEvents(ctx *gin.Context, params *operations.GetThingsDeviceNameEventsParams) *api.Response {
	panic("not implemented")
}

func (s *WebThingsApiService) GetThingsDeviceNameEventsEventType(ctx *gin.Context, params *operations.GetThingsDeviceNameEventsEventTypeParams) *api.Response {
	panic("not implemented")
}
