// cmd/blockchainnftvaultmanagerchainultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftvaultmanagerchainultra/internal/blockchainnftvaultmanagerchainultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftvaultmanagerchainultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
