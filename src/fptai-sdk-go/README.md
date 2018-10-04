# fptai-sdk-go
FPT.AI SDK for the Go programming language

## Usage
```
$ go get fptai-sdk-go
```
Example
```go
package main

import (
    "fptai-sdk-go"
)

func main() {
    client := fptai.NewClient("your_token")
    ...
}
```

## FPTAI CLI
SDK comes with a handy CLI tool that helps you import training data and evaluate.
```
$ export GOPATH=$(pwd)
$ export PATH=$PATH:$GOPATH/bin
$ go get fptai-sdk-go/cmd/fptai
$ fptai help
$ fptai train -t intent -i train.csv -token your_token
$ fptai test -t intent -i test.csv -token your_token
```

training.csv and test.csv file must be a CSV file and in following format:
```
intent_name1, intent utterance 1
intent_name1, intent utterance 2
intent_name2, intent utterance 3
intent_name1, intent utterance 4
...
```
