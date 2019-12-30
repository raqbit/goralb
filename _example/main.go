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

	var c color.RGBA

	if _, err := fmt.Sscanf(os.Args[1], "#%2x%2x%2x", &c.R, &c.G, &c.B); err != nil {
		panic(errors.New("could not parse input color"))
	}

	discoverTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	bm := goralb.NewBrushManager()
	defer bm.Close()

	fmt.Println("Turn on brush. This is also possible bluetooth-only using the mode select button")
	brush, err := bm.FindBrush(discoverTimeout)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found device, %+v\n", brush)

	err = brush.Connect()

	if err != nil {
		panic(err)
	}

	srTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = brush.WaitForServicesResolved(srTimeout); err != nil {
		panic(err)
	}

	if err = brush.SetColor(c); err != nil {
		panic(err)
	}

	brush.Disconnect()

	fmt.Println("Done!")
}
