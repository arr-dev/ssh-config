# ssh-config

Go app to parse [ssh_config(5)](http://linux.die.net/man/5/ssh_config) for given hosts.

## Installation

    go install github.com/arr-dev/ssh-config

## Usage

    ssh-config [options] HOSTNAME

    Options:
        -f, -file     SSH config file to parse, defaults to ~/.ssh/config
