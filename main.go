package main

import "fmt"

// Інтерфейси - це колекція сигнатур методів. Це контракт, який визначає, які методи повинен мати тип, але НЕ РЕАЛІЗУЄ ЇХ
//  type НазваІнтерфейсу interface{
// 	НазваМетоду(параметр) повертає
// 	ІншийМетод(параметр) повертає
//  }

//Duck type - "якщо щось ходить як качка і крякає як качка, то це качка!!!"
type Drawble interface {
	Draw()
}

type Resizeble interface {
	Resizeble(factor float64)
}

type GraphicElement interface {
	Drawble
	Resizeble
	GetInfo() string
}

type Square struct {
	Size float64
}

func (s *Square) Draw() {
	fmt.Printf("Draw square with size: %.f \n", s.Size)
}

func (s *Square) Resizeble(factor float64) {
	s.Size *= factor
	fmt.Printf("Draw square with size: %.f \n", s.Size)
}
func (s Square) GetInfo() string {
	return fmt.Sprintf("Square: sise=%.1f\n", s.Size)
}

func ProcessGraphic(g GraphicElement) {
	fmt.Println(g.GetInfo())
	g.Draw()
	g.Resizeble(1.5)
	g.Draw()
}
func main() {
	sq := &Square{Size: 10}
	ProcessGraphic(sq)
}
