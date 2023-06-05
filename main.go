package main
import (
	"fmt"
	"github.com/eiannone/keyboard"
	"klawiatura/pakiety"
	"log"
	"github.com/veandco/go-sdl2/sdl" 
)

func klawiatura (znak chan rune) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	} ()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		
		znak <- char
		
	}
}

func main() {
	swiat, err := wczytywanie.WczytajMapę("mapa.txt")
	if err != nil {
		log.Fatal (err)
	}

	dawajokno := make(chan *sdl.Window)
	
	go wczytywanie.RysujOkno(dawajokno)
	
	window := <- dawajokno

	jednostki, err := wczytywanie.WczytajMapę("jednostki.txt")
	if err != nil {
		log.Fatal (err)
	}
	
	for i, _ := range jednostki {
		for j, _ := range jednostki {
			wczytywanie.Rysuj(i, j, jednostki, swiat, window)
		}
	}
	znak := make(chan rune)
	go klawiatura(znak)
	
	gra := true
	pozostali := 0
	
	for gra {
		x := <- znak
		//fmt.Printf("%q\n", x)
		switch x { 
		case '\x00': //ESC
		
		
			return
		case 'm': //odkryj mape
			 wczytywanie.OdkryjMape(&swiat, window, jednostki)
		case 'w', 's', 'a', 'd'://W
			wczytywanie.PrzemiescGracza(x, &jednostki, window, swiat)
			wczytywanie.Widocznosc(jednostki, &swiat, window)
			wczytywanie.UsunSlady(&jednostki, swiat, window)
			wczytywanie.PrzemiescWroga(&jednostki, swiat, window)
			if wiezniowie := wczytywanie.IleWrogow(jednostki); wiezniowie != pozostali {
				pozostali = wiezniowie
				fmt.Println("W grze zostalo", pozostali,"wrogow")
				if pozostali == 0 {
					gra = false
					fmt.Println("Wszyscy zbiedzy zostali zapuszkowani!")
				}
			}
		}
	}
}

