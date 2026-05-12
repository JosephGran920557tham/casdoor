package controllers

import (
	"github.com/casdoor/casdoor/object"
)

// GetConfigInfo
// @Title GetConfigInfo
// @Tag Config API
// @Description get runtime configuration information
// @Success 200 {object} object.ConfigInfo
// @router /api/get-config-info [get]
func (c *ApiController) GetConfigInfo() {
	info := object.GetConfigInfo()
	c.Data["json"] = info
	c.ServeJSON()
}
