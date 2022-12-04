package calibratordate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gyu-young-park/lck_data_generator/crawler"
)

const CALIBRATION_DATE_FANDOM_FILE = "./calibrator-date-fandom-file.json"
const CALIBRATION_DATE_INVEN_FILE = "./calibrator-date-inven-file.json"

var calibrationDateFileMapper = map[string]string{
	string(crawler.CRAWLER_FANDOM_MODE): CALIBRATION_DATE_FANDOM_FILE,
	string(crawler.CRAWLER_INVEN_MODE):  CALIBRATION_DATE_INVEN_FILE,
}

type calibrateDateDataList struct {
	Data []calibrateDateData `json:"data"`
}

type calibrateDateData struct {
	VideoId string `json:"video-id"`
	Date    string `json:"date"`
}

var dateCalibrateTable = make(map[string]string)

func setDateCalibrateTable(crawlerMode string) {
	caliDataList, err := getCalibrateDate(crawlerMode)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, caliData := range caliDataList.Data {
		dateCalibrateTable[caliData.VideoId] = caliData.Date
	}
}

func getCalibrateDate(mode string) (*calibrateDateDataList, error) {
	path, isExist := calibrationDateFileMapper[mode]
	if !isExist {
		return nil, fmt.Errorf("getCalibrateDate Error: Not exist calibrator file in mode[%s]", mode)
	}
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var caliDataList calibrateDateDataList
	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteData, &caliDataList)
	if err != nil {
		return nil, err
	}
	return &caliDataList, nil
}
