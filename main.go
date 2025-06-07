package main

import "fmt"

// Інтерфейси - це колекція сигнатур методів. Це контракт, який визначає, які методи повинен мати тип, але НЕ РЕАЛІЗУЄ ЇХ
//  type НазваІнтерфейсу interface{
// 	НазваМетоду(параметр) повертає
// 	ІншийМетод(параметр) повертає
//  }

//Duck type - "якщо щось ходить як качка і крякає як качка, то це качка!!!"
type MusicInstrument interface {
	playMusic()
	tuneInstrument()
}

type Piano struct{}

func (p Piano) playMusic() {
	fmt.Println("Wasted! JuiceWRLD")
}

func (p Piano) tuneInstrument() {
	fmt.Println("tune... piano")
}

type Guitar struct{}

func (g Guitar) playMusic() {
	fmt.Println("Evil Jordan! Playboi Carti")
}

func (g Guitar) tuneInstrument() {
	fmt.Println("tune... guitar")
}

type Drums struct{}

func (d Drums) playMusic() {
	fmt.Println("liberation drums")
}

func (g Drums) tuneInstrument() {
	fmt.Println("tune... drums")
}

func playSongs(i MusicInstrument) {
	fmt.Println("getting ready")
	i.tuneInstrument()
	fmt.Println("play music")
	i.playMusic()
	fmt.Println("BB")
}

func main() {
	piano := Piano{}
	guitar := Guitar{}
	drum := Drums{}
	playSongs(piano)
	playSongs(guitar)
	playSongs(drum)
}
