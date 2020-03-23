# argparse
Simple and elegant command line argument parser for scripts that can be used as a shebang (Bash | Python | Perl etc.)

## Usage (shebang)
```
#!/usr/bin/argparse --executor=/bin/bash --arg "name::No Name Defined::Name"
```

## Usage (direct)
```
argparse --executor=/bin/bash --arg "name::No Name Defined::Name" printname.sh
```

The defined args would then be available as capitalized envirionment variables with the `ARG_<OPTION_NAME>` format.

## Example
```
○ cat printname.sh 
#!/Users/sreejith.kesavan/go/src/argparse/argparse --executor=/bin/bash --arg=name::NoNameDefined::Name

echo "This program prints names passed to the argument --name."
echo "Name: ${ARG_NAME}"

sreejith.kesavan at C02W91BSHV2R-blr in ~/go/src/argparse
○ ./printname.sh --name Sreejith
This program prints names passed to the argument --name.
Name: Sreejith

sreejith.kesavan at C02W91BSHV2R-blr in ~/go/src/argparse
○ ./printname.sh                
This program prints names passed to the argument --name.
Name: NoNameDefined

sreejith.kesavan at C02W91BSHV2R-blr in ~/go/src/argparse
○ ./printname.sh --help
Usage of ./printname.sh:
  -name string
    	Name (default "NoNameDefined")
```
