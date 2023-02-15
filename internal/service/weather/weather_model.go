package weather

type WeatherList struct {
	List []WeatherUnit `json:"list"`
}

type WeatherUnit struct {
	//Timestamp         time.Time          `json:"dt"`
	MainInfo          MainInfo           `json:"main"`
	WeatherDecription WeatherDescription `json:"weather"`
	Wind              Wind               `json:"wind"`
}
type Wind struct {
	Speed  float64 `json:"speed"` // Wind speed in meter/sec
	Degree float64 `json:"deg"`   // Wind direction, degrees (meteorological)
	Gust   float64 `json:"gust"`  // Wind gust in meter/sec
}

type WeatherDescription struct {
	Id          int    `json:"-"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"-"`
}

type MainInfo struct {
	Temperature   float64 `json:"temp"`
	FeelsLikeTemp float64 `json:"feels_like"`
	MinimumTemp   float64 `json:"temp_min"`
	MaximumTemp   float64 `json:"temp_max"`
	Humidity      float64 `json:"humidity"`
	Pressure      int     `json:"pressure"`
	SeaLevel      int     `json:"sea_level"`
	GroundLevel   int     `json:"grnd_level"`
}
