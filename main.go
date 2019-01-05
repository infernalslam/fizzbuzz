package main

import "fmt"

type dog struct {
	Name string
}

// AnimalName AnimalName
func AnimalName(d dog) string {
	return d.Name
}

func main() {
	d := dog{
		Name: "Two",
	}
	fmt.Println(d.Name)
}
