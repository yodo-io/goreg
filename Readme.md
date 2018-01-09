# Yet Another Docker Repository API Client

## Build

- Make sure [dep](https://github.com/golang/dep) is installed and up to date
- Run `make`

## Usage

### Help

```sh
./goreg help
```

### Private Registry (OSX Keychain)

```sh
REGISTRY_URL=registry.example.org

read REGISTRY_USERNAME REGISTRY_PASSWORD <<<$( \
  docker-credential-osxkeychain get <<<"$REGISTRY_URL" \
  | jq -r '"\(.Username) \(.Secret)"')

export REGISTRY_URL REGISTRY_USERNAME REGISTRY_PASSWORD
```
