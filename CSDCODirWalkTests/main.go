package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("CSDCO Directory walk filter testing")
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, VisitFile)
}

// VisitFile from https://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
func VisitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if fi.IsDir() {
		return nil // not a file.  ignore.
	}

	if caseInsenstiveContains(fp, "Geotek Data/whole-core data") {
		matched, err := filepath.Match("*_MSCL*", fi.Name())
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Println("GeoTek Whole Core:")
			fmt.Println(fp)
		}
	}

	if caseInsenstiveContains(fp, "Geotek Data/high-resolution MS data") {
		matched, err := filepath.Match("*.xls", fi.Name())
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			fmt.Println("GeoTek High Res:")
			fmt.Println(fp)
		}
	}

	return nil
}

func dirFunc(path string) (int, error) {
	var count int
	var FQP string

	err := filepath.Walk(path, func(_ string, f os.FileInfo, err error) error {

		FQP = fmt.Sprintf("%s/%s", path, f.Name())

		fmt.Println(f.Name())

		if !f.IsDir() {
			count = count + 1

			dir, file := filepath.Split(FQP)

			fmt.Printf("For path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)

			// set up our cases here to check (need an int to switch on though)..  based on these different actions will be done on the files
			// and different "types" will be set in the RDF graph for the file types

			// look for "Geotek Data"
			if caseInsenstiveContains(FQP, "whole-core data") {
				if caseInsenstiveContains(file, "_MSCL") {
					fmt.Printf("Geotek WholeCore data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
				}
			}

			// look for "ICD Data"  (add looking for .pdf on file name)
			if caseInsenstiveContains(FQP, "/ICD") { // use / to match to a directory...  avoid hitting "icd" in a directory name string
				fmt.Printf("ICD data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
			}

			// look for "CSD Data"
			// if caseInsenstiveContains(dir, "CSD Data") {
			// 	fmt.Printf("CSD data\nFor path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
			// }

			// fmt.Printf("For path:\t%s \nFor dir:\t%s\nFor file:\t%s \n\n", FQP, dir, file)
		}
		return err
	})

	return count, err
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
