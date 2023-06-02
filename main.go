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
		
 		// if key == keyboard.KeyEsc {
			// break
		// }
	}
}

func main() {
	swiat, err := wczytywanie.WczytajMapę("mapa.txt")
	if err != nil {
		log.Fatal (err)
	}

	// fmt.Println(swiat)
	dawajokno := make(chan *sdl.Window)
	
	go wczytywanie.RysujOkno(dawajokno)
	
	window := <- dawajokno
	
	// for i, _ := range swiat {
		// for j, _ := range swiat {
			// typ := swiat[i][j]
			// wczytywanie.RysujPole(i, j, typ, window)
		// }
	// }
	
	// jednostki := make([] [] int, 10)
	// for i := range jednostki {
		// jednostki [i] = make([]int, 10)
	// }
	
	jednostki, err := wczytywanie.WczytajMapę("jednostki.txt")
	if err != nil {
		log.Fatal (err)
	}
	
	for i, _ := range jednostki {
		for j, _ := range jednostki {
			// typ := jednostki[i][j]
			wczytywanie.Rysuj(i, j, jednostki, swiat, window)
		}
	}
	
	// fmt.Println(jednostki)
	
	znak := make(chan rune)
	go klawiatura(znak)
	
	gra := true
	
	for gra {
		x := <- znak
		fmt.Printf("%q\n", x)
		switch x { 
		case '\x00': //ESC
		
		
			return
		case 'w', 's', 'a', 'd'://W
			wczytywanie.PrzemiescGracza(x, &jednostki, window, swiat)
			wczytywanie.Widocznosc(jednostki, &swiat, window)
		}
	}
	
	//(znak rune, *jednostki int[][], window *sdl.Window, swiat int[][])

	
}

