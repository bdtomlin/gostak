package htm

import (
	"encoding/json"
	"fmt"
	"os"
)

var Assets = make(map[string]string)

func HashAssets() {
	jsonData, err := os.ReadFile("assetsManifest.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic("Must have a manifest.json file")
	}

	err = json.Unmarshal(jsonData, &Assets)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		panic("Must have a valid manifest.json file")
	}
}

func AssetPath(path string) string {
	return "/assets" + Assets[path]
}
