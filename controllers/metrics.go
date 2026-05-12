package controllers

import "github.com/casdoor/casdoor/object"

// GetMetrics godoc
// @Summary      Get runtime metrics
// @Description  Returns Go runtime and application metrics such as memory usage, goroutines, and uptime
// @Tags         metrics
// @Produce      json
// @Success      200  {object}  object.MetricsInfo
// @Router       /api/metrics [get]
func (c *ApiController) GetMetrics() {
	metrics := object.GetMetricsInfo()
	c.Data["json"] = metrics
	c.ServeJSON()
}
