package pkg

type Data struct {
	Latitude  float64   `json:"latitude,omitempty"`
	Longitude float64   `json:"longitude,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	Offset    float64   `json:"offset,omitempty"`
	Elevation int       `json:"elevation,omitempty"`
	Currently Currently `json:"currently,omitempty"`
	Minutely  Minutely  `json:"minutely,omitempty"`
	Hourly    Hourly    `json:"hourly,omitempty"`
	Daily     Daily     `json:"daily,omitempty"`
	Flags     Flags     `json:"flags,omitempty"`
}
type Currently struct {
	Time                 int     `json:"time,omitempty"`
	Summary              string  `json:"summary,omitempty"`
	Icon                 string  `json:"icon,omitempty"`
	NearestStormDistance int     `json:"nearestStormDistance,omitempty"`
	NearestStormBearing  int     `json:"nearestStormBearing,omitempty"`
	PrecipIntensity      float64 `json:"precipIntensity,omitempty"`
	PrecipProbability    float64 `json:"precipProbability,omitempty"`
	PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
	PrecipType           string  `json:"precipType,omitempty"`
	Temperature          float64 `json:"temperature,omitempty"`
	ApparentTemperature  float64 `json:"apparentTemperature,omitempty"`
	DewPoint             float64 `json:"dewPoint,omitempty"`
	Humidity             float64 `json:"humidity,omitempty"`
	Pressure             float64 `json:"pressure,omitempty"`
	WindSpeed            float64 `json:"windSpeed,omitempty"`
	WindGust             float64 `json:"windGust,omitempty"`
	WindBearing          float64 `json:"windBearing,omitempty"`
	CloudCover           float64 `json:"cloudCover,omitempty"`
	UvIndex              float64 `json:"uvIndex,omitempty"`
	Visibility           float64 `json:"visibility,omitempty"`
	Ozone                float64 `json:"ozone,omitempty"`
}
type MinutelyData struct {
	Time                 int     `json:"time,omitempty"`
	PrecipIntensity      float64 `json:"precipIntensity,omitempty"`
	PrecipProbability    float64 `json:"precipProbability,omitempty"`
	PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
	PrecipType           string  `json:"precipType,omitempty"`
}
type Minutely struct {
	Summary string         `json:"summary,omitempty"`
	Icon    string         `json:"icon,omitempty"`
	Data    []MinutelyData `json:"data,omitempty"`
}
type HourlyData struct {
	Time                 int     `json:"time,omitempty"`
	Icon                 string  `json:"icon,omitempty"`
	Summary              string  `json:"summary,omitempty"`
	PrecipIntensity      float64 `json:"precipIntensity,omitempty"`
	PrecipProbability    float64 `json:"precipProbability,omitempty"`
	PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
	PrecipAccumulation   float64 `json:"precipAccumulation,omitempty"`
	PrecipType           string  `json:"precipType,omitempty"`
	Temperature          float64 `json:"temperature,omitempty"`
	ApparentTemperature  float64 `json:"apparentTemperature,omitempty"`
	DewPoint             float64 `json:"dewPoint,omitempty"`
	Humidity             float64 `json:"humidity,omitempty"`
	Pressure             float64 `json:"pressure,omitempty"`
	WindSpeed            float64 `json:"windSpeed,omitempty"`
	WindGust             float64 `json:"windGust,omitempty"`
	WindBearing          float64 `json:"windBearing,omitempty"`
	CloudCover           float64 `json:"cloudCover,omitempty"`
	UvIndex              float64 `json:"uvIndex,omitempty"`
	Visibility           float64 `json:"visibility,omitempty"`
	Ozone                float64 `json:"ozone,omitempty"`
}
type Hourly struct {
	Summary string       `json:"summary,omitempty"`
	Icon    string       `json:"icon,omitempty"`
	Data    []HourlyData `json:"data,omitempty"`
}
type DailyData struct {
	Time                        int     `json:"time,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 int     `json:"sunriseTime,omitempty"`
	SunsetTime                  int     `json:"sunsetTime,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipAccumulation          float64 `json:"precipAccumulation,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int     `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int     `json:"temperatureLowTime,omitempty"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int     `json:"windGustTime,omitempty"`
	WindBearing                 float64 `json:"windBearing,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	UvIndex                     float64 `json:"uvIndex,omitempty"`
	UvIndexTime                 int     `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          int     `json:"temperatureMinTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          int     `json:"temperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime,omitempty"`
}
type Daily struct {
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
	Data    []DailyData `json:"data,omitempty"`
}
type SourceTimes struct {
	Gfs  string `json:"gfs,omitempty"`
	Gefs string `json:"gefs,omitempty"`
}
type Flags struct {
	Sources        []string    `json:"sources,omitempty"`
	SourceTimes    SourceTimes `json:"sourceTimes,omitempty"`
	NearestStation int         `json:"nearest-station,omitempty"`
	Units          string      `json:"units,omitempty"`
	Version        string      `json:"version,omitempty"`
}
