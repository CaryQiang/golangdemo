package worker

import (
	"net/http"
)

var MaxAlarmBufLen = 50

type AlarmServer struct {
	config     *Config
	httpServer *http.Server

	alarms   map[string]Alarm
	received chan bool
	stop     chan bool
	shutdown chan bool
}

func GetNewAlarmServer(conf *Config) (*AlarmServer, error) {
	alarm := AlarmServer{
		config: conf,
		//httpServer: &http.Server{Addr: conf.AlarmEndpoint},
	}

	return &alarm, nil
}

func (alarm *AlarmServer) Start() error {

	return nil
}

func (alarm *AlarmServer) Run() {

}

func (alarm *AlarmServer) Stop() {

}
