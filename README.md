# Magma Controller Rock

> **Warning**: This rock is under construction

Rock for Magma's orchestrator controller service built using 
[rockcraft](https://github.com/canonical/rockcraft). It provides a direct port of this 
[Dockerfile](https://github.com/magma/magma/blob/v1.6/orc8r/cloud/docker/controller/Dockerfile).

## Usage

```bash
docker pull ghcr.io/gruyaume/magma-orc8r-controller:1.6.1
docker run -it ghcr.io/gruyaume/magma-orc8r-controller:1.6.1
```

## TODO:
- Replace the build_context.py script with `organize`
- Replace all the files copying in the `override-build` with `organize`
- (maybe) separate files copying and go mod download to separate parts
- 