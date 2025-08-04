# Dawnshift

A screen shading tool for Hyprland that wraps hyprsunset and allows scheduling screen temperature changes.

## Build from source

##### Install Dependencies

```
yay -S go hyprsunset
```

##### Build

```
git clone git@github.com:rahuljjacob/dawnshift.git
cd dawnshift
go build -o dawnshift
install -Dm755 dawnshift /usr/bin/dawnshift
```

## Usage

```
Usage: dawnshift <command>
Commands:
  install  - Install systemd service
  apply    - Apply screen filter
  --help   - Show this help message
```

## Scheduling

Dawnshift uses a TOML configuration file to schedule screen temperature adjustments based on the time of day. The configuration file should be placed at

```
~/.config/dawnshift/dawnshift.toml
```

Dawnshift's configuration file allows you to schedule screen temperature adjustments based on the time of day. The default section sets the base temperature, while period entries define specific time ranges with different temperatures.

Example:

```
[default]
temperature = 6250

[[period]]
temperature = 6000
start_time = "18:30:00"
end_time = "20:00:00"

[[period]]
temperature = 5750
start_time = "20:00:00"
end_time = "22:00:00"

[[period]]
temperature = 5250
start_time = "22:00:00"
end_time = "01:00:00"

[[period]]
temperature = 4800
start_time = "01:00:00"
end_time = "06:00:00"
```

Install and enable the systemd timer/service units:

```
dawnshift install
systemctl --user enable --now dawnshift.timer
```

Add this to your hyprland config to ensure the correct screen temperature is set when you log in:

```
exec-once = dawnshift install
```
