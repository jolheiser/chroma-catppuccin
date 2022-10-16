package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
	_ "github.com/catppuccin/chroma-catppuccin"
)

//go:generate go run main.go
func main() {
	variants := []string{"mocha", "frappe", "macchiato", "latte"}

	for _, variant := range variants {
		if err := generateAsset(variant); err != nil {
			panic(err)
		}
	}
}

func generateAsset(variant string) error {
	src, err := os.ReadFile(fmt.Sprintf("../catppuccin-%s.go", variant))
	if err != nil {
		return err
	}

	fi, err := os.Create(variant + ".html")
	if err != nil {
		return err
	}
	defer fi.Close()

	if err := quick.Highlight(fi, string(src), "go", "html", fmt.Sprintf("catppuccin-%s", variant)); err != nil {
		return err
	}
	return nil
}
