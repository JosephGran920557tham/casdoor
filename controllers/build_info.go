package controllers

import (
	"github.com/casdoor/casdoor/object"
)

// GetBuildInfo
// @Title GetBuildInfo
// @Tag System API
// @Description get build information for the current deployment
// @Success 200 {object} object.BuildInfo
// @router /api/get-build-info [get]
func (c *ApiController) GetBuildInfo() {
	info := object.GetBuildInfo()
	c.ResponseOk(info)
}
