<!-- TOC ignore:true -->
# ⛅ go-weather

<!-- TOC ignore:true -->
## Weather forecasts in your system tray

> ⚠️ the project is still under development

## Table of contents

<!-- TOC -->

- [⛅ go-weather](#-go-weather)
    - [Table of contents](#table-of-contents)
        - [Requisites](#requisites)
        - [Installation](#installation)
        - [Development](#development)
            - [Use your icons](#use-your-icons)
        - [Screenshots](#screenshots)
        - [Credits](#credits)

<!-- /TOC -->

### Requisites

- A [Pirateweather](http://pirateweather.net/en/latest/) API Key (it's free!)
- Golang >= 1.16

### Installation

- Clone the repository
- Rename `config.yaml.dist` to `config.yaml`
- Add your data inside `config.yaml`

### Development

#### Use your icons

To use your icons follow these steps:
- Remove everything inside `icons` and `pack` folders, except for `Makefile` and `r2icon.go`
- Put your images inside `pack`
- Launch `make generate`
- Edit `r2icon.go` according to your new generated structs

### Screenshots

![macos](https://github.com/michelangelomo/go-weather/blob/master/screenshots/macos.png?raw=true)

![linux](https://github.com/michelangelomo/go-weather/blob/master/screenshots/linux.jpg?raw=true)

### Credits

Thanks to
- [lantern/systray](https://github.com/getlantern/systray) for the systray library
- [Makin-Things](https://github.com/Makin-Things/weather-icons) for the icons
- [cratonica](https://github.com/cratonica/2goarray) for the library to conver PNG to Go []byte struct