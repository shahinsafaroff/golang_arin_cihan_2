package main

import "fmt"

//maps are dictioniaries

func main() {
// First variant
	myMap := make(map[int]string) //[int] is the "key" and string is the "Value"
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)
// Second Variant
	airports := make(map[string]string)
	airports["IST"] = "Istambul"
	airports["ANK"] = "Ankara"
	airports["ANT"] = "Antalya"
	fmt.Println(airports)
	delete(airports, "ANT") //deletes item inside MAP
	fmt.Println(airports)
	airports["ERZ"] = "Erzurum" //adds item inside MAP

	for airport, city := range airports {
		fmt.Printf("%v: %v\n", airport, city) //formats a the map keys and values for looking good
	}

	//3. And now we will bypass city names inside maps and we will show only Airport names
	keys := make([]string, len(airports))
	//this means "keys" will be created as slice and slice items quantity will be derived from "airports" length
	i := 0
	for airport := range airports {
		keys[i] = airport
		i++
	}
	fmt.Println(keys)
	//4. And now we will bypass airport names inside maps and we will show only city names
	for i := range keys {
		fmt.Println(airports[keys[i]])
	}
}
