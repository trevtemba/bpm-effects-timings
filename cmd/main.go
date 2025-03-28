package main

import (
	"fmt"
	"os"
	"strconv"
)

type timings struct {
	verbDel  float64
	verbPre  float64
	compAtck float64
	compRel  float64
}

func main() {
	var bpm float64
	var time timings
	ok := false

	if len(os.Args) > 1 {
		var err error
		bpm, err = strconv.ParseFloat(os.Args[1], 32)

		if err != nil {
			bpm = 0
			fmt.Println("Command line arg must be bpm...")
		} else {
			ok = true
		}
	}

	for !ok {
		fmt.Printf("Enter bpm: ")
		_, err := fmt.Scan(&bpm)
		if err != nil {
			fmt.Println("Enter a valid number...")
		} else {
			ok = true
		}
	}

	baseNum := 60000 / bpm
	time.verbDel = baseNum * 4
	time.verbPre = baseNum / 16
	time.compAtck = baseNum / 16
	time.compRel = time.compAtck * 2

	fmt.Printf("\nEffects Timings:\n")
	fmt.Printf("----------------------------\n")
	fmt.Printf("%-20s %5.2fs\n", "Reverb Delay:", time.verbDel/1000)
	fmt.Printf("%-20s %5.2fms\n", "Reverb Pre-Delay:", time.verbPre)
	fmt.Printf("%-20s %5.2fms\n", "Compressor Attack:", time.compAtck)
	fmt.Printf("%-20s %5.2fms\n", "Compressor Release:", time.compRel)
	fmt.Printf("----------------------------\n")

}
