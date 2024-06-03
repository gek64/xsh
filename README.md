```
 ▄▄   ▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄   ▄▄▄▄▄▄ ▄▄   ▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄
█  █ █  █       █   ▄  █ █      █  █ █  █       █       █   █
█  █▄█  █▄▄▄▄   █  █ █ █ █  ▄   █  █▄█  █       █▄     ▄█   █
█       █▄▄▄▄█  █   █▄▄█▄█ █▄█  █       █     ▄▄█ █   █ █   █
█       █ ▄▄▄▄▄▄█    ▄▄  █      █▄     ▄█    █    █   █ █   █▄▄▄
 █     ██ █▄▄▄▄▄█   █  █ █  ▄   █ █   █ █    █▄▄  █   █ █       █
  █▄▄▄█ █▄▄▄▄▄▄▄█▄▄▄█  █▄█▄█ █▄▄█ █▄▄▄█ █▄▄▄▄▄▄▄█ █▄▄▄█ █▄▄▄▄▄▄▄█
```

- Make `v2ray` easier to install, uninstall and reload
- The application is installed in `/usr/local/bin/v2ray`, and the application data is stored
  in `/usr/local/etc/v2ray/` to avoid problems caused by future system changes.

## Usage

```
NAME:
   v2rayctl - v2ray quick install tool

USAGE:
   v2rayctl [global options] command [command options] 

VERSION:
   v3.01

COMMANDS:
   install, i  Install v2ray
   uninstall   Remove config,cache and uninstall v2ray
   update      Update v2ray
   reload      Reload service
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Install

```sh
# For example, the host is a linux kernel based system with amd64 architecture
# Download the app
curl -Lo /usr/local/bin/v2rayctl https://github.com/gek64/v2rayctl/releases/latest/download/v2rayctl-linux-amd64
# Give the app execute permission
chmod +x /usr/local/bin/v2rayctl
# Show help
/usr/local/bin/v2rayctl -h
```

## Compile

```sh
# Download application source code
git clone https://github.com/gek64/v2rayctl.git
# Compile the source code
cd v2rayctl
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details

## Credits

- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
