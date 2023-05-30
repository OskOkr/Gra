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
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		
		znak <- char
		
 		if key == keyboard.KeyEsc {
			break
		}
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
	
	for i, _ := range swiat {
		for j, _ := range swiat {
			typ := swiat[i][j]
			wczytywanie.RysujPole(i, j, typ, window)
		}
	}
	
	
	znak := make(chan rune)
	go klawiatura(znak)
	
	x := <- znak
	fmt.Printf("%q\n", x)
}

