package main

import (
   "github.com/ZoZo-182/gator/internal/config" 
   "fmt"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Errorf("Error with updated config %w", err)
    }
    cfg.SetUser("zianne")
    updatedcfg, err := config.Read()
    if err != nil {
        fmt.Errorf("Error with updated config %w", err)
    }
    fmt.Print(updatedcfg)
}
