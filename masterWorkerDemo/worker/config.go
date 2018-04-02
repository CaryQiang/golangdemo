package worker

import (
	"encoding/json"
	"golangdemo/masterWorkerDemo/pb"
	"io/ioutil"
)

type FaceRect struct {
	x int
	y int
	w int
	h int
}

//alarm
type Alarm struct {
	UUID        string
	Time        string
	VideoName   string
	VideoUUID   string
	LibraryName string
	LibraryUUID string
	LibraryType int
	CapImgURL   string
	CapFaceRect FaceRect
	TargetInfo  target.Target
}

type Config struct {
	AlarmEndpoint   string
	TargetEndpoint  string
	TargetGrpcpoint string
}

func LoadConfig(mf string) (*Config, error) {
	data, err := ioutil.ReadFile(mf)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
