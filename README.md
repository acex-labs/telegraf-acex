
# telegraf-acex

## Getting started

To build locally `go build -o ~/.local/bin/acex cmd/main.go`


To use within telegraf

```bash
cat << 'EOF' > acex.conf.toml
[[inputs.acex]]
  ## Optional URL if you plan to use HTTP later
  url = "https://localhost"
  insecure_skip_verify = true
EOF
```

Create telegraf config

```bash
cat << 'EOF' > telegraf.conf
[[inputs.execd]]
  ## One program to run as daemon.
  ## NOTE: process and each argument should each be their own string
  command = ["acex", "--config", "acex.conf.toml", "--poll_interval", "10s"]
EOF
```


Run telegraf with `telegraf --config telegraf.conf`




# TODO: CI

go build -ldflags "-X 'github.com/acex-labs/telegraf-acex/plugins/inputs/acex.Version={{Version}}'" -o ~/.local/bin/acex cmd/main.go