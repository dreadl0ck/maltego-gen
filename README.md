# maltego-gen

Generate Maltego Configurations with ease!

## Icons

Icons from the open source material design library can be automatically included as icons for your entities. 

List of all icons:

https://material.io/resources/icons/?style=twotone

This tool uses the twotone style by default.

### Colors 

You can generate colors via specifying the name from:

https://htmlcolorcodes.com/color-names/

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

to create a configuration to run your transforms locally with go (useful for debugging and development), use:

```yaml
executable: go
transforms:
  - id: LookupAddr
    input: maltego.IPv4Address
    description: Lookup Address
    args:
      - run
      - /Users/you/go/src/github.com/you/your-trx/cmd/transform/lookupAddr/main.go
  - id: ToDomainNames
    input: yourOrg.Domain
    description: To Domain Names
    args:
      - run
      - /Users/you/go/src/github.com/you/your-trx/cmd/transform/toDomains/main.go
```

## License

MIT