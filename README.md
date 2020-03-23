# argparse
Simple and elegant command line argument parser for scripts that can be used as a shebang (Bash | Python | Perl etc.)

## Usage (shebang)
```
#!/usr/bin/argparse --executor=/bin/bash --arg=name::NoNameDefined::Name
```

## Usage (direct)
```
argparse --executor=/bin/bash --arg=name::NoNameDefined::Name printname.sh
```

The defined args would then be available as capitalized envirionment variables with the `ARG_<OPTION_NAME>` format.

## Example
This example shows how `argparse` can be used as a `#!` in a `bash` script for commandline parsing. The following file is named `printname.sh`.

```
#!/usr/bin/argparse --executor=/bin/bash --arg=name::NoNameDefined::Name

echo "This program prints names passed to the argument --name."
echo "Name: ${ARG_NAME}"
```

```
$ ./printname.sh --name Sreejith
This program prints names passed to the argument --name.
Name: Sreejith
```

```
$ ./printname.sh                
This program prints names passed to the argument --name.
Name: NoNameDefined
```

```
$ ./printname.sh --help
Usage of ./printname.sh:
  -name string
    	Name (default "NoNameDefined")
```
