#!/bin/bash

wgo -file=.go -file=.templ -file="^assetsSrc/" -xfile=_templ.go -exit go tool templ generate \
  :: ./bin/tailwindcss -i ./input.css -o ./web/assetsSrc/css/styles.css \
  :: go run ./cmd/fingerprint/ \
  :: go run ./cmd/server/
