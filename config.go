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
	"github.com/dreadl0ck/maltego"
)

// YAML configuration
type config struct {

	// Organization name
	Org string `yaml:"org"`

	// Author name
	Author string `yaml:"author"`

	// Description of this Maltego configuration
	Description string `yaml:"description"`

	// Entities to generate
	Entities []*maltego.EntityCoreInfoExtended `yaml:"entities"`

	// Transforms to generate
	Transforms []*maltego.TransformCoreInfoExtended `yaml:"transforms"`

	// Executable file to use
	Executable string `yaml:"executable"`
}
