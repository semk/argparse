# argparse
Simple and elegant command line argument parser for scripts that can be used as a shebang (Bash | Python | Perl etc.)

## Installation
```
go get github.com/semk/argparse
cp $(go env GOPATH)/bin/argparse /usr/local/bin
```

## Usage
```
Usage of ./argparse:
  -arg value
    	Define custom argument to the program.
  -executor string
    	Path to the script executor (set this to shebang (#!)) (default "/bin/bash")
```

And the last arguments can be the script with the arguments to be called.

### in scripts #!
```
#!/usr/local/bin/argparse --executor=/bin/bash --arg=name::NoNameDefined::Name
```

### direct usage
```
argparse --executor=/bin/bash --arg=name::NoNameDefined::Name printname.sh --name Sreejith
```

The defined args would then be available as capitalized envirionment variables with the `ARG_<OPTION_NAME>` format.

## Example
This example shows how `argparse` can be used as a `#!` in a `bash` script for commandline parsing. The following file is named `printname.sh`.

```
#!/usr/local/bin/argparse --executor=/bin/bash --arg=name::NoNameDefined::Name

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
