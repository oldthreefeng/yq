## Usage

trans yaml to json, default is json

```
# use file
$ yq -f test.yaml  | jq .

# usr stdin

$ cat test.yaml | yq -f - | jq.

# use url

$ yq -f  https://fenghong.tech/yq/default.yaml   | jq .
```

trans json to yaml , output u should set to yaml. or will json to json.

```
# use file
$ yq -f test/test.json -o yaml  | jq .

# usr stdin

$ cat test/test.json  -o yaml | yq -f - | jq.

# use url

$ yq -f  https://fenghong.tech/yq/default.json -o yaml   | jq .
```