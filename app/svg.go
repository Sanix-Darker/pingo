package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// TODO: add more styling to this template
const TEMPLATE = `
	<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">
	  <text x="%d" y="%d" font-family="Verdana" font-size="20" fill="black">%d</text>
	</svg>
`

// CreateSVG
func CreateSVG(key string, number int) error {
	width, height, textX, textY := 200, 100, 50, 50
	svgContent := fmt.Sprintf(TEMPLATE, width, height, textX, textY, number)

	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Error: getting os.Getwd", err)
		return err
	}

	svgPath := path.Join(pwd, ".", fmt.Sprintf("./%s/%s.svg", ASSETS_PATH, key))
	if err := os.WriteFile(
		svgPath,
		[]byte(svgContent),
		0644,
	); err != nil {
		log.Fatalf("Error on Saving the asset svg %v", err)
		return err
	}

	return nil
}
