package reader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	InputAddr  string `json:"InputAddr"`
	OutputAddr string `json:"OutputAddr"`
	NumCore    int `json:"NumCore"`
	MaxGort    int `json:"MaxGort"`
	DataLength int `json:"DataLength"`
	CompressLev int `json:"CompressLev"`
	Parameters struct {
		MaxDistance    float64 `json:"MaxDistance "`
		MinScoreRANSAC float64 `json:"MinScoreRANSAC"`
		MinVertexPlane int `json:"MinVertexPlane"`
		MaxAnglePlane  float64 `json:"MaxAnglePlane"`
		MaxVertexQuit  int `json:"MaxVertexQuit"`
		MaxIteration   int `json:"MaxIteration"`
		NumBatch       int `json:"NumBatch"`
	} `json:"Parameters"`
}

func ReadConfig(filename string) Config {
	congfigFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer congfigFile.Close()

	byteValue, err := ioutil.ReadAll(congfigFile)
	if err != nil {
		log.Println(err)
	}

	var ConfigData Config
	err = json.Unmarshal(byteValue, &ConfigData)
	if err != nil {
		log.Println(err)
	}

	return ConfigData
}