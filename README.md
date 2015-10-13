# ssh-config

Go app to parse [ssh_config(5)](http://linux.die.net/man/5/ssh_config) for given hosts.

## Installation

    go install github.com/arr-dev/ssh-config

## Usage

    ssh-config [options] HOSTNAME

    Options:
      -f string
            Output format, p|plain, j|json, pretty (default "plain")
      -format string
            Output format, p|plain, j|json, pretty (default "plain")
      -file string
            SSH config file to parse (default "~/.ssh/config")
      -h    Output only Hostname from config
      -hostonly
            Output only Hostname from config

