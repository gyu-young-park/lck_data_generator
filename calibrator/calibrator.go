package calibrator

import (
	calibratordate "github.com/gyu-young-park/lck_data_generator/calibrator/date"
	"github.com/gyu-young-park/lck_data_generator/config"
)
type Calibrator struct {
	dataCalibrator *calibratordate.DateCalibrator
}

func NewCalibrator(config *config.Config) *Calibrator {
	return &Calibrator{
		dataCalibrator: calibratordate.NewDateCalibrator(config.CrawlerMode),
	}
}

func (c *Calibrator)GetCalibratedDate(videoId, date string) string {
	return c.dataCalibrator.Calibrate(videoId, date)
}