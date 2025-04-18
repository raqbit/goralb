package main

import (
	"context"
	"fmt"
	"github.com/Raqbit/goralb"
	"time"
)

func main() {
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

	fmt.Println("Subscribing to timer...")

	updates, err := brush.TimerUpdates()

	if err != nil {
		panic(err)
	}

	for upd := range updates {
		fmt.Printf("%d\n", upd)
	}

	fmt.Println("Disconnecting...")

	brush.Disconnect()

	fmt.Println("Done!")
}
