# cli

## Installation

1. Install [docker](https://docs.docker.com/engine/install/)

2. Install sath-engine and sath-cli:

```bash
wget https://download.sath.run/linux/sath.run && sudo bash sath.run
```

## Usage

Start sath-engine by:

```bash
sudo systemctl start sath
```

Run sath command:

```text
Usage:
  sath [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       Login
  logout      logout
  start       Start SATH engine
  status      Get job status
  stop        Stop SATH engine

Flags:
      --config string   config file (default is $HOME/.sath.yaml)
  -h, --help            help for sath

Use "sath [command] --help" for more information about a command.
```
