package uwsgi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Stat struct {
	Cwd               string `json:"cwd"`
	Gid               int    `json:"gid"`
	ListenQueue       int    `json:"listen_queue"`
	ListenQueueErrors int    `json:"listen_queue_errors"`
	Load              int    `json:"load"`
	Locks             []struct {
		User_0 int `json:"user 0"`
	} `json:"locks"`
	Pid         int `json:"pid"`
	SignalQueue int `json:"signal_queue"`
	Sockets     []struct {
		CanOffload int    `json:"can_offload"`
		MaxQueue   int    `json:"max_queue"`
		Name       string `json:"name"`
		Proto      string `json:"proto"`
		Queue      int    `json:"queue"`
		Shared     int    `json:"shared"`
	} `json:"sockets"`
	UID     int    `json:"uid"`
	Version string `json:"version"`
	Workers []struct {
		Accepting int `json:"accepting"`
		Apps      []struct {
			Chdir       string `json:"chdir"`
			Exceptions  int    `json:"exceptions"`
			ID          int    `json:"id"`
			Modifier1   int    `json:"modifier1"`
			Mountpoint  string `json:"mountpoint"`
			Requests    int    `json:"requests"`
			StartupTime int    `json:"startup_time"`
		} `json:"apps"`
		AvgRt int `json:"avg_rt"`
		Cores []struct {
			ID                int           `json:"id"`
			InRequest         int           `json:"in_request"`
			OffloadedRequests int           `json:"offloaded_requests"`
			ReadErrors        int           `json:"read_errors"`
			Requests          int           `json:"requests"`
			RoutedRequests    int           `json:"routed_requests"`
			StaticRequests    int           `json:"static_requests"`
			Vars              []interface{} `json:"vars"`
			WriteErrors       int           `json:"write_errors"`
		} `json:"cores"`
		DeltaRequests int    `json:"delta_requests"`
		Exceptions    int    `json:"exceptions"`
		HarakiriCount int    `json:"harakiri_count"`
		ID            int    `json:"id"`
		LastSpawn     int    `json:"last_spawn"`
		Pid           int    `json:"pid"`
		Requests      int    `json:"requests"`
		RespawnCount  int    `json:"respawn_count"`
		Rss           int    `json:"rss"`
		RunningTime   int    `json:"running_time"`
		SignalQueue   int    `json:"signal_queue"`
		Signals       int    `json:"signals"`
		Status        string `json:"status"`
		Tx            int    `json:"tx"`
		Vsz           int    `json:"vsz"`
	} `json:"workers"`
}

func GetStats(statAddr string) (*Stat, error) {
	client := http.Client{}
	fmt.Println(statAddr)
	req, err := http.NewRequest("GET", statAddr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	var stat = &Stat{}
	json.Unmarshal(buf, &stat)
	return stat, nil
}
