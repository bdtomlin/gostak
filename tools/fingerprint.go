package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var assets = make(map[string]string)

func Fingerprint() {
	jsonData, err := os.ReadFile("web/assetsManifest.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic("Must have a manifest.json file")
	}

	err = json.Unmarshal(jsonData, &assets)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		panic("Must have a valid web/manifest.json file")
	}

	err = os.RemoveAll("web/assets")
	if err != nil {
		panic(err)
	}

	for src, dest := range assets {
		src, dest = "web/assetsSrc"+src, "web/assets"+dest
		srcFile, err := os.Open(src)
		if err != nil {
			panic(fmt.Errorf("open src: %v", err))
		}
		defer srcFile.Close()

		dir := filepath.Dir(dest)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(fmt.Errorf("mkdir all: %v", err))
		}

		destFile, err := os.Create(dest)
		if err != nil {
			panic(fmt.Errorf("create destfile: %v", err))
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			fmt.Println("srcFile", src)
			fmt.Println("destFile", dest)
			panic(fmt.Errorf("io.copy: %v", err))
		}

		// Ensure the file has been written to disk
		err = destFile.Sync()
		if err != nil {
			panic(err)
		}
	}
}
