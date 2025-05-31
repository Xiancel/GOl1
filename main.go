package main

import "fmt"

type Car struct {
	Brand      string
	Model      string
	Years      int
	Price      float64
	IsElectric bool
}

type Garage struct {
	Cars []Car
}

func main() {
	car1 := Car{
		Brand:      "BMW",
		Model:      "M4",
		Years:      2024,
		Price:      80000,
		IsElectric: false,
	}
	car2 := Car{
		Brand:      "BMW",
		Model:      "X6M",
		Years:      2025,
		Price:      132100,
		IsElectric: false,
	}
	car3 := Car{
		Brand:      "Porsche",
		Model:      "Taycan turbo s",
		Years:      2023,
		Price:      168280,
		IsElectric: true,
	}

	garages := Garage{
		Cars: []Car{car1, car2, car3},
	}

	for i := 0; i < len(garages.Cars); i++ {
		fmt.Printf("Cars %d: %+v\n", i+1, garages.Cars[i])
	}
}
