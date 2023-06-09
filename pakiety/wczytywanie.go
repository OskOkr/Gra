package wczytywanie
import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func WczytajMapę (nazwaPliku string) ([] [] int, error) {
	file, err := os.Open (nazwaPliku)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	//tablica 40x40
	swiat := make([] [] int, 40)
	for i := range swiat {
		swiat [i] = make([]int, 40)
	}


	scanner := bufio.NewScanner(file)
	wiersz := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for kolumna, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			swiat [wiersz][kolumna] = num
		}
		wiersz++
	}


	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return swiat, nil
}