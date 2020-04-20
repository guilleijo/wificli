# wificli
MacOs wifi networks manager CLI in go


## Installation
```
go get github.com/guilleijo/wificli
```

## Usage

#### Help
Commands and flags information. Help information will be displayed if:
- No command used
- Help command or flag used
```
wificli
wificli help
wificli -h
```
![wificli](./assets/readme/help.gif)


#### List
Lists available wifi networks
```
wificli list
```
![wificli](./assets/readme/list.gif)

#### On/Off
Turn wifi on or off
```
wificli on
wificli off
```
![wificli](./assets/readme/on-off.gif)

#### Conn
Connect to wifi. Select from list or provide SSID and password using the `-n` and `-p` flags
```
# passing values directly
wificli conn -n <SSID> -p <password>

# if no flag is added you will be prompted to select a network and provide the password
wificli conn
```
![wificli](./assets/readme/conn-interactive.gif)
![wificli](./assets/readme/conn-flags.gif)

#### Speed
Download speed test
```
wificli speed
```
![wificli](./assets/readme/speed.gif)

#### Status
Current wifi connection status
```
wificli status
```
![wificli](./assets/readme/status.gif)

## Dependencies
- [cobra](https://github.com/spf13/cobra)
- [promptui](https://github.com/manifoldco/promptui)
- [spinner](https://github.com/briandowns/spinner)
- [go-fast](https://github.com/ddo/go-fast)

## TODO
- [x] Add gifs to README
- [ ] Support saved passwords
- [ ] Show wifi signal strength
- [ ] Implement speed test