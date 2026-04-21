package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

func main(){

	os.Setenv("WEATHER_API_KEY", "279a80b31a552206dfb3a6e9f71e0eaa")

	
	response, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=40.44&longitude=-79.99&current=temperature_2m")
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println(result);


}