package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 011、JSON movie
// go run 011_movie.go
// 输出：
// [{"title":"Casablanca","released":1942,"actors":["Humphrey Bogart","Ingrid Bergman"]},{"title":"Cool Hand Luke","released":1967,"color":true,"actors":["Paul Newman"]},{"title":"Bullitt","released":1968,"color":true,"actors":["Steve McQueen","Jacqueline Bisset"]}]
// [
//     {
//         "title": "Casablanca",
//         "released": 1942,
//         "actors": [
//             "Humphrey Bogart",
//             "Ingrid Bergman"
//         ]
//     },
//     {
//         "title": "Cool Hand Luke",
//         "released": 1967,
//         "color": true,
//         "actors": [
//             "Paul Newman"
//         ]
//     },
//     {
//         "title": "Bullitt",
//         "released": 1968,
//         "color": true,
//         "actors": [
//             "Steve McQueen",
//             "Jacqueline Bisset"
//         ]
//     }
// ]
// [{Casablanca} {Cool Hand Luke} {Bullitt}]
// [{Title:Casablanca Year:1942 Color:false Actors:[Humphrey Bogart Ingrid Bergman]} {Title:Cool Hand Luke Year:1967 Color:true Actors:[Paul Newman]} {Title:Bullitt Year:1968 Color:true Actors:[Steve McQueen Jacqueline Bisset]}]
// Casablanca 1942 false [Humphrey Bogart Ingrid Bergman]
// Cool Hand Luke 1967 true [Paul Newman]
// Bullitt 1968 true [Steve McQueen Jacqueline Bisset]
func main() {
	// JavaScript 对象表示法（JSON）是一种发送和接收结构化信息的标准协议
	type Movie struct {
		Title  string   `json:"title"`
		Year   int      `json:"released"`
		Color  bool     `json:"color,omitempty"` // omitempty 忽略空值字段，不忽略会返回当前类型零值，忽略则字段不返回
		Actors []string `json:"actors"`
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	// 输出：
	// [{"title":"Casablanca","released":1942,"actors":["Humphrey Bogart","Ingrid Bergman"]},{"title":"Cool Hand Luke","released":1967,"color":true,"actors":["Paul Newman"]},{"title":"Bullitt","released":1968,"color":true,"actors":["Steve McQueen","Jacqueline Bisset"]}]
	// 这种紧凑的形式虽然包含了全部的信息，但是很难阅读
	// json.MarshalIndent 函数将产生整齐缩进的输出：
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	// 输出：
	// [
	// 	{
	// 		"title": "Casablanca",
	// 		"released": 1942,
	// 		"actors": [
	// 			"Humphrey Bogart",
	// 			"Ingrid Bergman"
	// 		]
	// 	},
	// 	{
	// 		"title": "Cool Hand Luke",
	// 		"released": 1967,
	// 		"color": true,
	// 		"actors": [
	// 			"Paul Newman"
	// 		]
	// 	},
	// 	{
	// 		"title": "Bullitt",
	// 		"released": 1968,
	// 		"color": true,
	// 		"actors": [
	// 			"Steve McQueen",
	// 			"Jacqueline Bisset"
	// 		]
	// 	}
	// ]

	// 解码，可以只解需要的字段
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed: %s", err)
	}
	fmt.Println(titles)
	// 输出：
	// [{Casablanca} {Cool Hand Luke} {Bullitt}]

	var decode []Movie
	if err := json.Unmarshal(data, &decode); err != nil {
		log.Fatalf("JSON unmarshal failed: %s", err)
	}
	fmt.Printf("%+v\n", decode)
	// 输出：
	// [{Title:Casablanca Year:1942 Color:false Actors:[Humphrey Bogart Ingrid Bergman]} {Title:Cool Hand Luke Year:1967 Color:true Actors:[Paul Newman]} {Title:Bullitt Year:1968 Color:true Actors:[Steve McQueen Jacqueline Bisset]}]
	for _, v := range decode {
		fmt.Println(v.Title, v.Year, v.Color, v.Actors)
	}
	// 输出：
	// Casablanca 1942 false [Humphrey Bogart Ingrid Bergman]
	// Cool Hand Luke 1967 true [Paul Newman]
	// Bullitt 1968 true [Steve McQueen Jacqueline Bisset]
}
