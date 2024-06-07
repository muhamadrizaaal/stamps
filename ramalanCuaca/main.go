package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "sort"
    "time"
)

const (
    apiKey     = "2019809906bdbc89c310fb7485320e3f" // Menggunakan API key yang diberikan
    city       = "Jakarta"
    apiURL     = "http://api.openweathermap.org/data/2.5/forecast"
    units      = "metric"
    dateFormat = "Mon, 02 Jan 2006"
)

type Forecast struct {
    List []struct {
        Dt   int64 `json:"dt"`
        Main struct {
            Temp float64 `json:"temp"`
        } `json:"main"`
        DtTxt string `json:"dt_txt"`
    } `json:"list"`
}

func main() {
    url := fmt.Sprintf("%s?q=%s&units=%s&appid=%s", apiURL, city, units, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Printf("Failed to fetch data: %s\n", resp.Status)
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("Response Body:", string(body))
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var forecast Forecast
    if err := json.Unmarshal(body, &forecast); err != nil {
        panic(err)
    }

    dailyTemps := make(map[string][]float64)
    for _, entry := range forecast.List {
        date := time.Unix(entry.Dt, 0).Format("2006-01-02")
        dailyTemps[date] = append(dailyTemps[date], entry.Main.Temp)
    }

    var dates []string
    for date := range dailyTemps {
        dates = append(dates, date)
    }

    sort.Strings(dates)

    fmt.Println("Weather Forecast:")
    for _, date := range dates {
        temps := dailyTemps[date]
        avgTemp := average(temps)
        parsedDate, _ := time.Parse("2006-01-02", date)
        fmt.Printf("%s: %.2f°C\n", parsedDate.Format(dateFormat), avgTemp)
    }
}

func average(nums []float64) float64 {
    var sum float64
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}
