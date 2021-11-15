package main

import "fmt"

func main() {

	map1 := make(map[int]float64)

	map1[1] = 1.0

	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	for key := range map1 {
		fmt.Printf("key is: %v \n", key)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

}
