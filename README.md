# ngos

#### (EXPERIMENTAL, there are still many bugs)

Tool for comparing two `CSV` files, and create the output of that difference

TODO:
- Implement Concurrency
- Modular Code
- Handling `CSV` file with multiple column

### Usage

- Build binary from source

```shell
$ go get github.com/Bhinneka/ngos
$ go install github.com/Bhinneka/ngos/cmd/ngos
$ ngos -v
```

```shell
$ ngos -old a.csv -new b.csv -o output.csv
```