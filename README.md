# maltego-gen

[![Go Report Card](https://goreportcard.com/badge/github.com/dreadl0ck/maltego)](https://goreportcard.com/report/github.com/dreadl0ck/maltego)
[![License](https://img.shields.io/badge/license-MIT-green)](https://raw.githubusercontent.com/dreadl0ck/maltego/master/LICENSE)

Generate [Maltego](https://maltego.com) Configurations with ease!

## Install

Install the _maltego-gen_ commandline tool:

    go install github.com/dreadl0ck/maltego-gen

## Icons

Icons from the open source material design library can be automatically included as icons for your entities. 

See the list of all icons [here](https://material.io/resources/icons/?style=twotone).

This tool uses the _twotone_ style by default.

### Colors 

You can generate colors via specifying a lower-case [color name](https://htmlcolorcodes.com/color-names) or via hex code.

## Usage

Create a YAML file to describe your Entities and Transforms:

```yaml
org: YourOrg
author: Your Name
description: What this config is used for

entities:
  # choose the display for the new entity
  - name: Domain
    # set the image for your entity
    image:
      # choose image from the material icon list
      name: domain
      # choose a color or hex code for the svg (if empty, will default to black)
      color: black
    description: A domain
    # set parent entity
    parent: maltego.Domain
    fields:
      - name: unicode
        description: Unicode representation of domain name
      - name: ascii
        description: ASCII representation of domain name

# local transforms
transforms:
  # example transformation that invokes a binary file: simply pass the path to the binary
  - id: LookupAddr
    input: maltego.IPv4Address
    description: Lookup Address
    executable: /path/to/your/binary
  
  # example transformation that invokes a script:
  - id: ToDomainNames
    input: dittotrx.IDNDomain
    description: To Domain Names
    executable: /path/to/your/interpreter
    args:
      - /path/to/your/script
      - -debug
```

to create a configuration to recompile your Go transforms on every run (useful for debugging and development), use:

```yaml
executable: go
workingDir: ~/go/src/github.com/you/your-trx

transforms:
  - id: LookupAddr
    input: maltego.IPv4Address
    description: Lookup Address
    args:
      - run
      - cmd/transform/lookupAddr/main.go
  - id: ToDomainNames
    input: yourOrg.Domain
    description: To Domain Names
    args:
      - run
      - cmd/transform/toDomains/main.go
```

Put the YAML into a file and invoking maltego-gen with it will generate the following Maltego configuration:

```
$ maltego-gen config.yml 
material icon repository exists, pulling
bootstrapped configuration archive for Maltego
packing maltego yourorg archive
packed maltego yourorg archive
copied generated file to /Users/you/YourOrg.mtz

$ tree ./yourorg
├── Entities
│   └── yourorg.Domain.entity
├── EntityCategories
│   └── yourorg.category
├── Icons
│   └── yourorg
│       ├── domain_black.svg
│       ├── domain_black.xml
│       ├── domain_black24.svg
│       ├── domain_black32.svg
│       ├── domain_black48.svg
│       └── domain_black96.svg
├── Servers
│   └── Local.tas
├── TransformRepositories
│   └── Local
│       ├── yourorg.LookupAddr.transform
│       ├── yourorg.LookupAddr.transformsettings
│       ├── yourorg.ToDomainNames.transform
│       └── yourorg.ToDomainNames.transformsettings
├── TransformSets
│   └── YourOrg.set
└── version.properties

8 directories, 15 files
$ du -h yourorg.mtz
8.0K    yourorg.mtz
```

> If the config is named _maltego.yml_ you can simply invoke maltego-gen in the same directory, and there is no need to specify the file as an argument. 

On the first execution, the material icon repository will be cloned to the _/tmp_ folder.
Subsequent runs will only check if the repo is up-to-date.

The resulting _yourorg.mtz_ file can be imported into Maltego.

## License

MIT