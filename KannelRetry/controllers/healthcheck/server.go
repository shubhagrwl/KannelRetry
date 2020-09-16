package healthcheck

import (
	"encoding/json"
	"net/http"

	"github.com/KannelRetry/response"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//Organization -
type Organization struct {
	gorm.Model
	ID         string
	Email      string
	Phone      string
	EmployeeNo string
	OrgID      string
}

// HealthCheck ...
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Healthy = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	logrus.WithFields(logrus.Fields{"api": "Healthcheck api"}).Info("Runs successfuly")
	return
}
