package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
        return arr[1:]
    } else {
        return arr[:len(arr)-1]
    }
}

func main() {
	array1 := []interface{}{1, 2, 3, 4, 5}
    fmt.Println(removeLeftRight(array1, "left")) // [2 3 4 5]

    array2 := []interface{}{1, 2, 3, 4, 5}
    fmt.Println(removeLeftRight(array2, "right")) // [1 2 3 4]

    array3 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    fmt.Println(removeLeftRight(array3, "left")) // [Budi Joko Tono]

    array4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    fmt.Println(removeLeftRight(array4, "right")) // [Edo Budi Joko]
}
