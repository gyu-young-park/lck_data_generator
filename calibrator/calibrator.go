package calibrator

import (
	calibratordate "github.com/gyu-young-park/lck_data_generator/calibrator/date"
	calibratorvideo "github.com/gyu-young-park/lck_data_generator/calibrator/video"
	"github.com/gyu-young-park/lck_data_generator/config"
	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

type Calibrator struct {
	dataCalibrator calibratordate.DateCalibrator
	videoInjector  calibratorvideo.VideoInjector
}

func NewCalibrator(config *config.Config) *Calibrator {
	return &Calibrator{
		dataCalibrator: calibratordate.NewDateCalibrator(config.CrawlerMode),
		videoInjector:  calibratorvideo.NewVideoInjectorImpl(config.CrawlerMode),
	}
}

func (c *Calibrator) GetCalibratedDate(videoId, date string) string {
	return c.dataCalibrator.Calibrate(videoId, date)
}

func (c *Calibrator) SetOmittedVideoInVideoMapper(videoMapper videoitem.VideoItemListMapper, videoService videostatistics.Service) {
	c.videoInjector.Calibrate(videoMapper, videoService)
}
