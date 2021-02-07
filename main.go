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
	"github.com/davecgh/go-spew/spew"
	"github.com/dreadl0ck/maltego"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	flag.Parse()

	data, err := ioutil.ReadFile(*flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	var c = new(config)
	err = yaml.UnmarshalStrict(data, c)
	if err != nil {
		log.Fatal(err)
	}

	c.Executable, err = exec.LookPath(c.Executable)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(c)

	var (
		org    = strings.Title(c.Org)
		ident  = strings.ToLower(org)
		prefix = ident + "."
		//machinePrefix = ident + "_"
		propsPrefix = "properties."
	)

	maltego.GenMaltegoArchive(ident, org)
	for _, e := range c.Entities {
		for i, f := range e.Fields {
			if f.Nullable {
				e.Fields[i] = maltego.NewStringField(f.Name, f.Description)
			} else {
				e.Fields[i] = maltego.NewRequiredStringField(f.Name, f.Description)
			}
		}
		maltego.GenEntity(
			org,
			ident,
			prefix,
			propsPrefix,
			ident,
			e.Name,
			e.Icon,
			e.Description,
			e.Parent,
			"black",
			nil,
			e.Fields...,
		)
	}

	for _, t := range c.Transforms {
		//args := []string{"run", t.ID}
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
			true,
		)
	}

	simpleTransforms := []*maltego.TransformCoreInfo{}
	for _, t := range c.Transforms {
		simpleTransforms = append(simpleTransforms, &maltego.TransformCoreInfo{
			ID:          t.ID,
			InputEntity: t.InputEntity,
			Description: t.Description,
		})
	}
	maltego.GenServerListing(prefix, ident, simpleTransforms)
	maltego.GenTransformSet(c.Org, c.Description, prefix, ident, simpleTransforms)
	maltego.PackMaltegoArchive(ident)

	file := c.Org + ".mtz"
	path := filepath.Join(os.Getenv("HOME"), file)
	maltego.CopyFile(file, path)
	fmt.Println("copied generated file to", path)
}
