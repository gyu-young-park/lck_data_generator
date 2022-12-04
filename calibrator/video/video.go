package calibratorvideo

import (
	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

type VideoInjector interface {
	Calibrate(videoitem.VideoItemListMapper, videostatistics.Service)
}
