package icons

var MAPPING = map[string][]byte{
	"clear":               ClearDay,
	"clear-day":           ClearDay,
	"clear-night":         ClearNight,
	"rain":                Rainy1,
	"snow":                Snowy1,
	"sleet":               RainAndSleetMix,
	"wind":                Wind,
	"fog":                 Fog,
	"cloudy":              Cloudy,
	"partly-cloudy-day":   Cloudy1Day,
	"partly-cloudy-night": Cloudy1Night,
}

// This method convert the status obtained from the provider to an appropriate icon
// Possible values are: clear-day, clear-night, rain, snow, sleet, wind, fog, cloudy, partly-cloudy-day, partly-cloudy-night
func Convert(status string) []byte {
	return MAPPING[status]
}
