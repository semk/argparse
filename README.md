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
