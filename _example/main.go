package main

import (
	"context"
	"fmt"
	"github.com/Raqbit/goralb"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	advert, err := goralb.FindBrush(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", advert)

	fmt.Println("Done!")
}
