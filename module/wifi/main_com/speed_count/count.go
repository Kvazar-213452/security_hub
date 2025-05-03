package speed_count

import (
	"github.com/showwin/speedtest-go/speedtest"
)

func Start_speed_count() (float64, float64) {
	client := speedtest.New()

	serverList, err := client.FetchServers()
	if err != nil {
		return 0.0, 0.0
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return 0.0, 0.0
	}

	err = targets[0].DownloadTest()
	if err != nil {
		return 0.0, 0.0
	}

	downloadSpeed := (float64(targets[0].DLSpeed) * 8) / 1e6

	err = targets[0].UploadTest()
	if err != nil {
		return 0.0, 0.0
	}

	uploadSpeed := (float64(targets[0].ULSpeed) * 8) / 1e6

	return downloadSpeed, uploadSpeed
}
