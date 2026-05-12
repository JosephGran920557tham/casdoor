package controllers

import (
	"github.com/casdoor/casdoor/object"
)

type HealthController struct {
	BaseController
}

// GetHealthStatus godoc
// @Title GetHealthStatus
// @Tag Health API
// @Description get the health status of the server
// @Success 200 {object} object.HealthStatus
// @router /health [get]
func (c *HealthController) GetHealthStatus() {
	status := object.GetHealthStatus()

	if status.Status != "ok" {
		c.Ctx.ResponseWriter.WriteHeader(503)
	}

	c.Data["json"] = status
	c.ServeJSON()
}

// Ping godoc
// @Title Ping
// @Tag Health API
// @Description ping the server to check if it is alive
// @Success 200 {string} string "pong"
// @router /ping [get]
func (c *HealthController) Ping() {
	c.Data["json"] = map[string]string{"message": "pong"}
	c.ServeJSON()
}
