package calibratordate

import "fmt"

type DateCalibratorImpl struct {
}

func NewDateCalibrator(crawlerMode string) *DateCalibratorImpl {
	setDateCalibrateTable(crawlerMode)
	return &DateCalibratorImpl{}
}

func (d *DateCalibratorImpl) Calibrate(videoId string, date string) string {
	calibratedDate, isExist := dateCalibrateTable[videoId]
	fmt.Println("cal:", calibratedDate)
	fmt.Println("video:", videoId)
	if !isExist {
		return date
	}
	return calibratedDate
}
