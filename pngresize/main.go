// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  ./pngresize/resize.go copy from codestack 
    convert png from xhdpi -> hdpi -> mdpi
*/

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

var (
	xhdpi = "./xhdpi/"
	hdpi  = "./hdpi/"
	mdpi  = "./mdpi/"

	toDdpi = true
)

func main() {

	// open input file
	fi, err := os.Open("xhdpi")
	if err != nil {
		panic(err)
	}

	dirinfo, err := fi.Readdir(0)
	for _, dir := range dirinfo {
		skip := dir.IsDir()
		if !skip {
			//convert(dir, 20, 15)
			fmt.Println(dir.Name())
			convert(dir.Name(), true)
			convert(dir.Name(), false)
		}
	}

}

func convert(file string, toDdpi bool) {

	infile, err := os.Open(xhdpi + file)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := infile.Close(); err != nil {
			panic(err)
		}
	}()

	pic, err := png.Decode(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, ": %v\n", err)
	}
	b := pic.Bounds()
	var outImg image.Image
	if toDdpi {
		outImg = Resize(pic, b, b.Dx()*20/15, b.Dy()*20/15)
	} else {
		outImg = Resize(pic, b, b.Dx()/2, b.Dy()/2)
	}

	newName := ""
	if toDdpi {
		newName = hdpi + file
	} else {
		newName = mdpi + file
	}
	// open output file
	fo, err := os.Create(newName)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	png.Encode(fo, outImg)
}
