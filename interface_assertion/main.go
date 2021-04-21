package main

import "fmt"

func main() {
	//fmt.Println("Program is running...")

	// var a interface{} = 5
	// var b interface{} = 6.5
	// var c interface{} = "AA"
	// var d interface{} = [3]int{1, 2, 3}

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// var data interface{} = map[string]interface{}{
	// 	"person": map[string]string{
	// 		"name":  "Zahid",
	// 		"city":  "Dhaka",
	// 		"hobby": "something",
	// 	},
	// }

	// datacity := data.(map[string]interface{})["person"].(map[string]string)["city"]
	// fmt.Println(datacity)

	var data interface{} = map[string]interface{}{
		"username":            "nahidhasan98",
		"name":                "Nahid Hasan",
		"honor":               303,
		"clan":                "",
		"leaderboardPosition": 111698,
		"skills":              [3]int{1, 2, 3},
		"ranks": map[string]interface{}{
			"overall": map[string]interface{}{
				"rank":  -5,
				"name":  "5 kyu",
				"color": "yellow",
				"score": 297,
			},
			"languages": map[string]interface{}{
				"go": map[string]interface{}{
					"rank":  -5,
					"name":  "5 kyu",
					"color": "yellow",
					"score": 297,
				},
				"cpp": map[string]interface{}{
					"rank":  -8,
					"name":  "8 kyu",
					"color": "white",
					"score": 2,
				},
			},
		},
		"codeChallenges": map[string]interface{}{
			"totalAuthored":  0,
			"totalCompleted": 32,
		},
	}

	datarank := data.(map[string]interface{})["ranks"].(map[string]interface{})["languages"].(map[string]interface{})["go"].(map[string]interface{})["score"]
	fmt.Println(datarank)

	// var e interface{} = map[string]interface{}{
	// 	"first":     10,
	// 	"second":    20.7,
	// 	"something": "Nahid",
	// 	"other": map[string]interface{}{
	// 		"a": 4,
	// 		"b": 8,
	// 	},
	// }
	// var eNew = e.(map[string]interface{})["other"].(map[string]int)["a"]
	// fmt.Println(eNew)

	// var e1 = e.(map[string]interface{})
	// var e2 = e1["other"]
	// var n1 = e2.(map[string]int)["a"]

	// fmt.Println(n1)

	// var eNew = e.(map[string]interface{})
	// fmt.Println(eNew["something"])

	// mp := e.(map[string]int)["first"]
	// fmt.Println(mp)

	// f := map[string]int{
	// 	"first":  10,
	// 	"second": 20,
	// }
	// fmt.Println(f["first"])
}
