# EasyWeather
An easy-to-use web application for retrieving the real-time weather of a user-inputted location, written in Golang. Developed using [Gin](https://github.com/gin-gonic/gin) web framework and weather data derived from [OpenWeatherMap](https://github.com/briandowns/openweathermap) API. Returns basic data like temperature, wind speed, UV index with severity and forecast for the next 5 days.

## Configuration
1. Obtain an API key from [OpenWeatherMap](https://openweathermap.org/) and signing up.
2. Set API key as an environment variable or update it in the ```config.json``` file.
3. Make sure the relevant dependencies like Gin are set up in the system.

## API Calls
- A call is made using the ```NewCurrent()``` function which return a ```CurrentWeatherData``` pointer.

  ```
  w, err := owm.NewCurrent("C", "EN", apikey)
  err = w.CurrentByName(location)
  ```

- UV index information is retrieved using ```http.Get()``` where the endpoint is supplied with the API key, latitude and longitude.

  ```
  uvIndexResp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%f&lon=%f", apikey, w.GeoPos.Latitude, w.GeoPos.Longitude))
  ```
## Planned Features
- A ```Chart.js``` graph populated with the forecasted values.
- An interactive map with the requested location.
- Animation and background images based on the weather conditions.

## Known Issues
- Favicon does not load.(Fixed)
- External linking of ```CSS``` file does not load.(Fixed)

## Documentation Links
- [OpenWeatherMap](https://pkg.go.dev/github.com/briandowns/openweathermap#section-readme)
- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- [Golang Official Documentation](https://go.dev/doc/)

