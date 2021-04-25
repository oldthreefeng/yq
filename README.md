## Usage

trans yaml to json, default is yaml

```
# use file
$ yq -f test.yaml -o json | jq .

# usr stdin

$ cat test.yaml | yq -f - -o json| jq .

# use url

$ yq -f  https://fenghong.tech/yq/default.yaml -o json  | jq .
```

trans json to yaml.

```
# use file
$ yq -f test/test.json  

# usr stdin

$ cat test/test.json | yq -f - 

# use url

$ yq -f  https://fenghong.tech/yq/default.json 
```


## docker use

docker only support url if u don't use volume
 
```
$ docker run --rm  louisehong/yq  -f https://fenghong.tech/yq/default.yaml -o json  | jq .
```

## use goreleaser and github actions to cicd
