# cli

## Installation

```bash
wget https://download.sath.run/linux/sath.run && sudo bash sath.run
```

## Usage

Install and start [docker](https://docs.docker.com/engine/install/)

Start sath-engine by:

```bash
sudo systemctl start sath
```

Run sath command:

```text
A CLI tool to run algorithm for SATH platform in your local machine,
Complete documentation is available at https://github.com/sath-run/cli

Usage:
  sath [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  jobs        List jobs
  login       Login
  logout      logout
  start       Start SATH engine
  status      Get sath engine status
  stop        Stop SATH engine

Flags:
      --config string   config file (default is $HOME/.sath.yaml)
  -h, --help            help for sath

Use "sath [command] --help" for more information about a command.
```
