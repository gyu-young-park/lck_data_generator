package calibratorvideo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gyu-young-park/lck_data_generator/crawler"
)

const CALIBRATION_OMITTED_VIDEO_FANDOM_FILE = "./calibrator-omiited-video-fandom-file.json"
const CALIBRATION_OMITTED_VIDEO_INVEN_FILE = "./calibrator-omiited-video-inven-file.json"

var calibrationOmittedVideoFileMapper = map[string]string{
	string(crawler.CRAWLER_FANDOM_MODE): CALIBRATION_OMITTED_VIDEO_FANDOM_FILE,
	string(crawler.CRAWLER_INVEN_MODE):  CALIBRATION_OMITTED_VIDEO_INVEN_FILE,
}

type calibrateOmittedVideoDataList struct {
	Data []calibrateOmittedVideoData `json:"data"`
}

type calibrateOmittedVideoData struct {
	PlayListTitle string   `json:"playlist-title"`
	Season        string   `json:"season"`
	VideoIdList   []string `json:"video-id-list"`
	Date          string   `json:"date"`
}

// date-videoid
var omittedVideoCalibrateTable = make(map[string]*calibrateOmittedVideoData)

func setOmittedVideoCalibrateTable(crawlerMode string) {
	caliVideoDataList, err := getCalibrateOmittedVideo(crawlerMode)
	if err != nil {
		fmt.Println("[setOmittedVideoCalibrateTable]Err:", err)
		return
	}
	for _, caliData := range caliVideoDataList.Data {
		omittedVideoCalibrateTable[caliData.Date] = &caliData
	}
}

func getCalibrateOmittedVideo(mode string) (*calibrateOmittedVideoDataList, error) {
	path, isExist := calibrationOmittedVideoFileMapper[mode]
	if !isExist {
		return nil, fmt.Errorf("getCalibrateDate Error: Not exist calibrator file in mode[%s]", mode)
	}
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var caliVideoDataList calibrateOmittedVideoDataList
	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteData, &caliVideoDataList)
	if err != nil {
		return nil, err
	}
	return &caliVideoDataList, nil
}
