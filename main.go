package main

import "fmt"

// Інтерфейси - це колекція сигнатур методів. Це контракт, який визначає, які методи повинен мати тип, але НЕ РЕАЛІЗУЄ ЇХ
//  type НазваІнтерфейсу interface{
// 	НазваМетоду(параметр) повертає
// 	ІншийМетод(параметр) повертає
//  }

//Duck type - "якщо щось ходить як качка і крякає як качка, то це качка!!!"
type Animal interface {
	MakeSound() string
}
type Dog struct {
	Name  string
	Breed string
}

type Bird struct {
	Name    string
	Species string
}

func (d Dog) MakeSound() string {
	return "woof!"
}

func (b Bird) MakeSound() string {
	return "tweet!"
}

// func PrintDogSound(d Dog) {
// 	fmt.Println(d.MakeSound())
// }

// func PrintBirdSound(b Bird) {
// 	fmt.Println(b.MakeSound())
// }

func PrintAminalSound(a Animal) {
	fmt.Println(a.MakeSound())
}
func main() {
	bird := Bird{Name: "Kesha", Species: "Parrot"}
	dog := Dog{Name: "", Breed: ""}
	PrintAminalSound(bird)
	PrintAminalSound(dog)
}
