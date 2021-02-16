/*
 * MALTEGO-GEN - Commandline tool to generate Maltego configurations.
 * Copyright (c) 2021 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/dreadl0ck/maltego"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	icongen "github.com/dreadl0ck/material-icon-gen"
)

func main() {

	flag.Parse()

	// read config file
	data, err := ioutil.ReadFile(*flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal config
	var c = new(config)
	err = yaml.UnmarshalStrict(data, c)
	if err != nil {
		log.Fatal(err)
	}

	// resolve executable path
	c.Executable, err = exec.LookPath(c.Executable)
	if err != nil {
		log.Fatal(err)
	}

	var (
		org         = strings.Title(c.Org)
		ident       = strings.ToLower(org)
		prefix      = ident + "."
		propsPrefix = "properties."
	)

	maltegoSizes := []int{16, 24, 32, 48, 96}

	// image name to colors
	var coloredIcons = map[string][]string{}

	for _, e := range c.Entities {
		if e.Image != nil {

			// use black as default color if none has been specified
			if e.Image.Color == "" {
				e.Image.Color = "black"
			}

			// collect color for image name
			if _, ok := coloredIcons[e.Image.Name]; !ok {
				coloredIcons[e.Image.Name] = []string{e.Image.Color}
			} else {
				coloredIcons[e.Image.Name] = append(coloredIcons[e.Image.Name], e.Image.Color)
			}
		}
	}

	if len(coloredIcons) > 0 {

		// generate icons
		icongen.GenerateIconsSVG(
			*flagImagePath,
			icongen.DefaultSvgURL,
			maltegoSizes,
			coloredIcons,
			func(newBase string, color string) {
				maltego.CreateXMLIconFile(newBase + "_" + color)
			},
		)
	}

	// prepare archive
	maltego.GenMaltegoArchive(ident, org)

	// generate entities
	for _, e := range c.Entities {
		for i, f := range e.Fields {
			if f.Nullable {
				e.Fields[i] = maltego.NewStringField(f.Name, f.Description)
			} else {
				e.Fields[i] = maltego.NewRequiredStringField(f.Name, f.Description)
			}
		}

		maltego.GenEntity(
			*flagImagePath,
			org,
			ident,
			prefix,
			propsPrefix,
			ident,
			e.Name,
			e.Image.Name,
			e.Description,
			e.Parent,
			e.Image.Color,
			nil,
			e.Fields...,
		)
	}

	// generate transforms
	for _, t := range c.Transforms {
		maltego.GenTransform(
			c.Org,
			c.Author,
			prefix,
			ident,
			t.ID,
			t.Description,
			t.InputEntity,
			c.Executable,
			t.Args,
			*flagTransformDebug,
		)
	}

	// collect as simpleTransforms, to pass it to the server listing generation
	simpleTransforms := []*maltego.TransformCoreInfo{}
	for _, t := range c.Transforms {
		simpleTransforms = append(simpleTransforms, &maltego.TransformCoreInfo{
			ID:          t.ID,
			InputEntity: t.InputEntity,
			Description: t.Description,
		})
	}

	// generate a listing to include the local transforms
	maltego.GenServerListing(prefix, ident, simpleTransforms)

	// add transform set
	maltego.GenTransformSet(c.Org, c.Description, prefix, ident, simpleTransforms)

	// pack archive
	maltego.PackMaltegoArchive(ident)

	// copy the archive into the home directory
	if *flagCopyToHomeDir {
		file := c.Org + ".mtz"
		path := filepath.Join(os.Getenv("HOME"), file)
		maltego.CopyFile(file, path)
		fmt.Println("copied generated file to", path)
	}
}
