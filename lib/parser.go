package parser

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type GeoLocationBranches struct {
	id          int
	branch      string
	geoLocation string
}

func ReadCsvFile(filePath string) (map[string]*GeoLocationBranches, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	geoLocation := map[string]*GeoLocationBranches{}
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return geoLocation, err
		}

		g := &GeoLocationBranches{}
		if g.id, err = strconv.Atoi(row[0]); err != nil {
			return nil, err
		}
		g.branch = row[1]
		g.geoLocation = row[2]
		geoLocation[row[0]] = g
	}
}
