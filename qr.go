package main

import (
	"os"

	"github.com/skip2/go-qrcode"
)

var RLevel = qrcode.High

func Create(content string) ([][]bool, error) {
	qr, err := qrcode.New(content, RLevel)
	if err != nil {
		return nil, err
	}

	return qr.Bitmap(), nil
}

// Draw draws the bitmap to the terminal, using a single whitespace as bit character
// and the black / white escape sequence at every bit change
func Draw(qr [][]bool, black, white string) {
	os.Stdout.WriteString(white)
	last := false // white = false, black = true

	for _, row := range qr {
		for _, bit := range row {
			if bit != last {
				if bit {
					os.Stdout.WriteString(black)
				} else {
					os.Stdout.WriteString(white)
				}
				last = bit
			}

			os.Stdout.Write([]byte{' '})
		}
		os.Stdout.Write([]byte{'\n'})
	}
}
