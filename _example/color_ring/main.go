package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Raqbit/goralb"
	"image/color"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("please give hexcode as argument")
	}

	var newColor color.RGBA

	if _, err := fmt.Sscanf(os.Args[1], "#%2x%2x%2x", &newColor.R, &newColor.G, &newColor.B); err != nil {
		panic(errors.New("could not parse input color_ring"))
	}

	discoverTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	bs, err := goralb.NewScanner()

	if err != nil {
		panic(fmt.Errorf("could not create brush scanner: %w", err))
	}

	defer bs.Close()

	fmt.Println("Turn on brush. This is also possible bluetooth-only using the mode select button")
	brush, err := bs.FindBrush(discoverTimeout)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found device, %s\n", brush)

	connectTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("Connecting...")

	if err = brush.Connect(connectTimeout); err != nil {
		panic(err)
	}

	fmt.Println("Retrieving current color_ring...")

	prevColor, err := brush.GetColor()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Current color_ring: (R: %d, G: %d, B: %d)\n", prevColor.R, prevColor.G, prevColor.B)
	fmt.Printf("New color_ring: (R: %d, G: %d, B: %d)\n", newColor.R, newColor.G, newColor.B)

	fmt.Println("Setting new color_ring...")

	if err = brush.SetColor(newColor); err != nil {
		panic(err)
	}

	fmt.Println("Turning on...")

	if err = brush.SetRingEnabled(true); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("Turning off...")

	if err = brush.SetRingEnabled(false); err != nil {
		panic(err)
	}

	fmt.Println("Disconnecting...")

	brush.Disconnect()

	fmt.Println("Done!")
}
