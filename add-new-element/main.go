package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
        return append([]int{newData}, data...)
    } else if position == "down" {
        return append(data, newData)
    }
    return data
}

func main() {
	data := []int{1, 2, 3, 4, 5}
    newData := 6
    position := "up"

    // menambahkan element baru di awal array
    result := AddElement(data, newData, position)
    fmt.Println(result) // output: [6 1 2 3 4 5]

    position = "down"

    // menambahkan element baru di akhir array
    result = AddElement(data, newData, position)
    fmt.Println(result) // output: [1 2 3 4 5 6]
}
