package object

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type HealthStatus struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
}

func GetHealthStatus() *HealthStatus {
	status := &HealthStatus{
		Status:    "ok",
		Timestamp: time.Now(),
		Checks:    make(map[string]string),
	}

	if err := checkDatabaseHealth(); err != nil {
		status.Status = "degraded"
		status.Checks["database"] = err.Error()
	} else {
		status.Checks["database"] = "ok"
	}

	return status
}

func checkDatabaseHealth() error {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT 1").Exec()
	return err
}
