package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/michelangelomo/go-weather/icons"
	"github.com/michelangelomo/go-weather/pkg"
	log "github.com/sirupsen/logrus"
)

var (
	config    Config
	provider  pkg.Provider
	menuItems map[string]map[string]*systray.MenuItem
)

func main() {
	// Config stuff
	err := config.parseConfig()
	if err != nil {
		log.Fatalf("Configuration parsing error: %v", err)
	}
	err = config.isValid()
	if err != nil {
		log.Fatalf("Configuration validation error: %v", err)
	}

	// Init weather provider
	provider.NewProvider(config.APIKey)

	menuItems = map[string]map[string]*systray.MenuItem{}
	for _, l := range config.getLocations() {
		menuItems[l.Name] = make(map[string]*systray.MenuItem)
	}

	// Start systray
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("go-weather")
	systray.SetTemplateIcon(icons.ClearDay, icons.ClearDay)

	systray.SetTooltip("")
	mQuit := systray.AddMenuItem("Quit", "Exit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icons.Wind)
	systray.AddSeparator()

	showMainLocation()
	showOtherLocations()
	startWatcher()
}

func showMainLocation() {
	// Get main location
	location := config.getMainLocation()
	req := pkg.Request{
		Provider: &provider,
		Unit:     config.Weather.Unit,
		Lat:      location.Lat,
		Lon:      location.Lon,
	}
	req.Get()
	if req.Error != nil {
		log.Warnf("Error retrieving %s weather: %v", location.Name, req.Error)
		systray.SetTitle("API Error")
		systray.SetTemplateIcon(icons.Hurricane, icons.Hurricane)
	} else {
		systray.SetTitle(fmt.Sprintf("%s %v °C", location.Name, req.Data.Currently.Temperature))
		icon := icons.Convert(req.Data.Currently.Icon)
		systray.SetTemplateIcon(icon, icon)
	}

	showLocationInMenu(location.Name, *req.Data)
}

func showOtherLocations() {
	locations := config.getOtherLocations()
	for _, l := range locations {
		req := pkg.Request{
			Provider: &provider,
			Unit:     config.Weather.Unit,
			Lat:      l.Lat,
			Lon:      l.Lon,
		}
		req.Get()
		if req.Error != nil {
			log.Warnf("Error retrieving %s weather: %v", l.Name, req.Error)
		} else {
			showLocationInMenu(l.Name, *req.Data)
		}
	}
}

func startWatcher() {
	interval := parseInterval(config.Weather.Interval)
	log.Infof("Start watcher with interval of %s", interval.String())
	for range time.Tick(interval) {
		log.Info("Refreshing...")
		go showMainLocation()
		go showOtherLocations()
	}
}

func showLocationInMenu(name string, data pkg.Data) {
	t := fmt.Sprintf("%s %v °C", name, data.Currently.Temperature)

	loc := createOrUpdateItem(name, "title", t, data.Currently.Icon, false, nil)

	// Main informations
	createOrUpdateItem(name, "humidity", fmt.Sprintf("Humidity %v %%", data.Currently.Humidity*100), "", true, nil)
	createOrUpdateItem(name, "wind", fmt.Sprintf("Wind %v / %v", data.Currently.WindSpeed, data.Currently.WindBearing), "", true, nil)

	// Sub menu informations
	createOrUpdateItem(name, "windgust", fmt.Sprintf("Wind gust %v", data.Currently.WindGust), "wind", true, loc)
	createOrUpdateItem(name, "minutely", fmt.Sprintf("Minutely: %s", data.Minutely.Summary), data.Minutely.Icon, true, loc)
	createOrUpdateItem(name, "hourly", fmt.Sprintf("Hourly: %s", data.Hourly.Summary), data.Hourly.Icon, true, loc)
	createOrUpdateItem(name, "daily", fmt.Sprintf("Daily: %s", data.Daily.Summary), data.Daily.Icon, true, loc)

	systray.AddSeparator()
}

func createOrUpdateItem(name, key, value, iconName string, isDisabled bool, parentItem *systray.MenuItem) *systray.MenuItem {
	// Check if menu item is already instantiated
	_, ok := menuItems[name][key]
	// It doens't exists
	if !ok {
		if parentItem != nil {
			menuItems[name][key] = parentItem.AddSubMenuItem(name, "")
		} else {
			menuItems[name][key] = systray.AddMenuItem(name, "")
		}
	}
	menuItems[name][key].SetTitle(value)
	if isDisabled {
		menuItems[name][key].Disable()
	}

	if iconName != "" {
		icon := icons.Convert(iconName)
		if len(icon) != 0 {
			menuItems[name][key].SetTemplateIcon(icon, icon)
		}
	}

	return menuItems[name][key]
}

func parseInterval(s string) (t time.Duration) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		t = time.Duration(n) * time.Second
	} else {
		t, err = time.ParseDuration(s)
	}
	if err != nil {
		t = time.Duration(0)
	}
	if t < 0 {
		t = -t
	}
	return t
}

func onExit() {
}
