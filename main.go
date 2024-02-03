package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/serpapi/serpapi-golang"
)

func main() {
	auth := map[string]string{
		"api_key": "6ac8ca94f6989ebd667711160bd2a1a53c98da4952741fdea5368c6aadcf8337",
	}
	client := serpapi.NewClient(auth)

	items := []string{
		"allen wrenches", "gerbil feeders", "toilet seats", "electric heaters", "trash compactors",
		"juice extractor", "shower rods", "water meters", "walkie-talkies", "copper wires",
		"safety goggles", "radial tires", "BB pellets", "rubber mallets", "fans",
		"dehumidifiers", "picture hangers", "paper cutters", "waffle irons", "window shutters",
		"paint removers", "window louvres", "masking tape", "plastic gutters", "kitchen faucets",
		"folding tables", "weather stripping", "jumper cables", "hooks and tackle", "grout and spackle",
		"power foggers", "spoons and ladles", "pesticides for fumigation", "high-performance lubrication",
		"metal roofing", "water proofing", "multi-purpose insulation", "air compressors", "brass connectors",
		"wrecking chisels", "smoke detectors", "tire guages", "hamster cages", "thermostats",
		"bug deflectors", "trailer hitch demagnetizers", "automatic circumcisers", "tennis rackets",
		"angle brackets", "Duracells and Energizers", "soffit panels", "circuit brakers", "vacuum cleaners",
		"coffee makers", "calculators", "generators", "matching salt and pepper shakers",
	}

	for index, item := range items {
		parameter := map[string]string{
			"engine": "google_images",
			"q":      item,
		}
		rsp, err := client.Search(parameter)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rsp["search_metadata"].(map[string]interface{})["status"] != "Success" {
			fmt.Println(rsp)
			return
		}

		if rsp["images_results"] == nil {
			fmt.Println("no images results")
			return
		}

		results := rsp["images_results"].([]interface{})
		if len(results) > 0 {
			resultMap := results[0].(map[string]interface{})
			if original, ok := resultMap["original"].(string); ok {
				downloadImage(original, fmt.Sprintf("%d_%s.jpg", index+1, strings.ReplaceAll(item, " ", "_")))
				// break // Only download the first image for each item
			}
		}
	}
}

func downloadImage(url, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error saving file %s: %v\n", filename, err)
		return
	}

	fmt.Printf("Downloaded %s\n", filename)
}
