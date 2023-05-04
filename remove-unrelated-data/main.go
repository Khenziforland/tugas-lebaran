package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
    return dataMap
}

func main() {
	data1 := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	fmt.Println(removeUnrelated(data1, "address"))

	data2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	fmt.Println(removeUnrelated(data2, "jump"))

	data3 := map[string]interface{}{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	fmt.Println(removeUnrelated(data3, "tiga"))
}
