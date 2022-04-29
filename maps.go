package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.3}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	} else {
		fmt.Printf("%s is passing!.\n", name)
	}
}

func main() {
	//var myMap[string]float64

	//var ranks map[string]int
	//ranks = make(map[string]float64)

	/*ranks := make(map[string]float64)
	/*ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3*/
	/*ranks := map[string]float64{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])*/
	
	/*elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"*/
	elements := map[string]string{
		"H": "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
	
	/*isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true*/
	/*isPrime := map[int]bool{
		4: false,
		7: true,
	}
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])*/

	myMap := map[string]float64{
		"a": 1.2, 
		"b": 5.6,
	}
	fmt.Println(myMap["a"])
	fmt.Println(myMap["b"])
	fmt.Println(myMap)
	fmt.Printf("%#v\n", myMap)
	
	//emptyMap := map[string]float64{}
	
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])
	
	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I haven't been assigned"])
	
	/*counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])*/
	
	var nilMap map[int]string
	fmt.Printf("%#v\n", nilMap)
	//nilMap[3] = "three"
	
	var notNilMap map[int]string = make(map[int]string)
	fmt.Printf("%#v\n", notNilMap)
	notNilMap[3] = "three"
	
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	_, ok = counters["b"]
	fmt.Println(ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	
	status("Alma")
	status("Rohit")
	status("Carl")
	
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
	
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {		// map[string]int{"a": 2, "c": 1, "e": 2}
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
	// a: 2
	// b: not found
	// c: 1
	// d: not found
	// e: 2
	
	//var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	
	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}