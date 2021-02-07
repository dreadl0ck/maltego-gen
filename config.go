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

type config struct {
	Org         string                               `yaml:"org"`
	Author      string                               `yaml:"author"`
	Description string                               `yaml:"description"`
	Entities    []*maltego.EntityCoreInfo            `yaml:"entities"`
	Transforms  []*maltego.TransformCoreInfoExtended `yaml:"transforms"`
	Executable  string                               `yaml:"executable"`
}
