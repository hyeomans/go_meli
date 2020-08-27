Mercadolibre Go SDK
-----

## ¿Por qué?

Me dí cuenta que la versión que sugieren en el sitio de ML, no toma en cuenta `context`, ni tampoco
agrega todos los endpoints.

## ¿Cómo usarla?

Este es un WIP (work in progress). Sólo he agregado lós endpoints públicos:

```
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	meli "github.com/hyeomans/go_meli"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	log.Printf("run : Started")
	defer log.Println("run : Completed")
	http := &http.Client{}

	meliConfig := meli.Config{
		HTTP:        http,
	}
	meliClient, err := meli.New(meliConfig)

	if err != nil {
		return fmt.Errorf("could not initialize meliClient: %s", err)
	}

	r, err := meliClient.Sites.ListingPrices(ctx, "MLM", 30000)
	if err != nil {
		return err
	}
	fmt.Println(r)

	return nil
}
```