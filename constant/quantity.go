package constant

type QuantityData struct {
	Resolution string
	Scale      string
}

var QUANTITY_MAP map[string]QuantityData = map[string]QuantityData{
	string(QUEUE_MP4_360_P): {
		Resolution: "360p",
		Scale:      "640:360",
	},
	string(QUEUE_MP4_480_P): {
		Resolution: "480p",
		Scale:      "854:480",
	},
	string(QUEUE_MP4_720_P): {
		Resolution: "720p",
		Scale:      "1280:720",
	},
	string(QUEUE_MP4_1080_P): {
		Resolution: "1080p",
		Scale:      "1920:1080",
	},
}
