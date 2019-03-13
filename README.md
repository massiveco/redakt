# redakt

Strip sensitive content out of YAML files

## Example Usage

> kustomize build | redakt --kind Namespace | kubectl diff