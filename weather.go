package main 

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"

	"net/http"

	"fmt"

	"io/ioutil"

	"encoding/json"

	"fyne.io/fyne/v2/canvas"

	"image/color"
)

func showWeatherApp(w fyne.Window) {
	a:= app.New();
	w := a.NewWindow("Weather App");
	w.Resize(fyne.NewSize(500,500));

	// API Part
	res , err:= http.Get("http://api.openweathermap.org/data/2.5/weather?q=gurgaon&APPID=cadcd1596d3a10d00b58d34d39a98b94")

	if err!=nil{
		fmt.Println(err);

	}
	defer res.Body.Close();

	body, err:= ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err);
	}

	weather , err:=UnmarshalWeather(body)
	if err!=nil{
		fmt.Println(err);
	}

	image:= canvas.NewImageFromFile("weather.jpg");
	image.FillMode = canvas.ImageFillOriginal

	label_1 := canvas.NewText("Weather Details", color.White)
	label_1.TextStyle = fyne.TextStyle{Bold:true}

	label_2 := canvas.NewText(fmt.Sprintf("Country %s", weather.Sys.Country),color.White)

	label_3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather.Wind.Speed),color.White)

	label_4 := canvas.NewText(fmt.Sprintf("Temperature %2f", weather.Main.Temp),color.White)

	weatherContainer:=container.NewVBox(
			label_1,
			image,
			label_2,
			label_3,
			label_4,
			container.NewGridWithColumns(1),
		),
	)


w.SetContent(container.NewBorder(panelContent,nil,nil,nil,weatherContainer),)

w.ShowAndRun();
}
// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()





func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
