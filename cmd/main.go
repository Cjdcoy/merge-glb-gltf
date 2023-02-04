package main

import (
	"github.com/qmuntal/gltf"
	log "github.com/sirupsen/logrus"
	"merge-glb-gltf/pkg/gltfMerger"
	"os"
)

func main() {
	var err error
	var merger *gltfMerger.GltfMerger
	var cube gltf.Document
	var plane gltf.Document
	var f *os.File

	if f, err = os.Open("data/cone.gltf"); err != nil {
		log.Fatal(err)
	}
	if err = gltf.NewDecoder(f).Decode(&plane); err != nil {
		log.Fatal(err)
	}
	merger = gltfMerger.NewGltfMerger(&plane)
	f.Close()
	if f, err = os.Open("data/cylinder.gltf"); err != nil {
		log.Fatal(err)
	}
	if err = gltf.NewDecoder(f).Decode(&cube); err != nil {
		log.Fatal(err)
	}
	merger.Merge(&cube)
	merger.WriteDoc("data/test.gltf")
}
