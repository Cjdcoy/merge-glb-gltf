package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"polyhouse/pkg/objHandler"
)

type (
	objPlexer struct {
		objs []*objHandler.Object
		nbVertex int64
		nbTexture int64
		nbNormal int64
	}
)
func writeObj(object *objHandler.Object) {
	for _, face := range object.Faces {
		for _, point := range face.Points {
			fmt.Println(*point.Texture, "/", *point.Vertex, "/", *point.Normal)
		}
	}
}

func (e *objPlexer) updateFacesIndex(obj *objHandler.Object) {
	for _, face := range obj.Faces {
		for _, point := range face.Points {
			point.Vertex.Index += e.nbVertex
			point.Texture.Index += e.nbTexture
			point.Normal.Index += e.nbNormal
		}
	}
	e.nbVertex += int64(len(obj.Vertices))
	e.nbTexture += int64(len(obj.Textures))
	e.nbNormal += int64(len(obj.Normals))
}

func (e *objPlexer) addObj(obj *objHandler.Object) {
	if obj != nil {
		e.updateFacesIndex(obj)
		e.objs = append(e.objs, obj)
	} else {
		log.Error("error is nil")
	}
}

func (e *objPlexer) writeObjs() {
	f, err := os.OpenFile("data/test.objHandler", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Panic(err)
	}
	os.Open("data/test.objHandler")
	writer := objHandler.NewWriter(f)
	for  _, obj := range e.objs {
		writer.Write(obj)
	}
}

func loadObj(path string) (*objHandler.Object, error) {
	reader, _ := os.Open(path)
	objReader := objHandler.NewReader(reader)
	return objReader.Read()
}

type (
	MeasurementLevel struct {
		// Name of the level
		Name string `yaml:"name" json:"name"`
		// Range of the level
		Range []int `yaml:"range" json:"range"`
	}

	MeasurementValue struct {
		// Label is the name (ex: R1V)
		Label string `yaml:"label" json:"label"`
		// Value is the read measurement
		Value int `yaml:"value" json:"value"`
	}

	MeasurementData struct {
		Name string `yaml:"name" json:"name"`
		// Min full-scale value low
		Min int `yaml:"min" json:"min"`
		// Max full-scale value max
		Max int `yaml:"max" json:"max"`
		//
		Levels []MeasurementLevel `yaml:"levels" json:"levels"`
		// Unity is the unity for the measurement (example: °C for the temperature)
		Unity string `yaml:"unity" json:"unity"`
		// Measurement list
		Measurement []MeasurementValue `yaml:"measurement" json:"measurement"`
	}

	Measurement struct {
		Data *MeasurementData
	}
)