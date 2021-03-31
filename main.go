package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var inputFile string
	var protoDir string
	var xxxTags string
	var inputFiles []string
	flag.StringVar(&inputFile, "input", "", "path to input file")
	flag.StringVar(&protoDir, "proto_dir", "", "path to proto dir")
	flag.StringVar(&xxxTags, "XXX_skip", "", "skip tags to inject on XXX fields")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}
	if len(inputFile) != 0 {
		inputFiles = append(inputFiles, inputFile)
	}
	if len(protoDir) != 0 {
		pbGoFiles, err := getPbGoFileByProtoDir(protoDir)
		if err != nil {
			log.Fatal(err)
		}
		inputFiles = append(inputFiles, pbGoFiles...)
	}
	if len(inputFiles) == 0 {
		log.Fatal("input file or proto dir is mandatory")
	}
	for _, file := range inputFiles {
		areas, err := parseFile(file, xxxSkipSlice)
		if err != nil {
			log.Fatal(err)
		}
		if err = writeFile(file, areas); err != nil {
			log.Fatal(err)
		}
	}

}