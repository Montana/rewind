package main

import (
	"image/gif"
	"os"
)

func main() {

	in, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	gif, err := gif.DecodeAll(in)
	if err != nil {
		panic(err)
	}

	for i := len(gif.Image) - 1; i >= 0; i-- {
		gif.Image = append(gif.Image, gif.Image[i])
		gif.Delay = append(gif.Delay, gif.Delay[i])
	}

	gif.Disposal = nil

	err = gif.EncodeAll(out, gif)
	if err != nil {
		panic(err)
	}

}
