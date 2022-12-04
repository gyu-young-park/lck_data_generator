package calibratordate

import "fmt"

type DateCalibrator struct {
}

func NewDateCalibrator(crawlerMode string) *DateCalibrator{
	setDateCalibrateTable(crawlerMode)
	return &DateCalibrator{}
}

func (d *DateCalibrator)Calibrate(videoId string, date string) string {
	calibratedDate, isExist := dateCalibrateTable[videoId]
	fmt.Println("cal:",calibratedDate)
	fmt.Println("video:",videoId)
	if !isExist {
		return date
	}
	return calibratedDate
}