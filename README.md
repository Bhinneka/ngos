# ngos

#### (EXPERIMENTAL, there are still many bugs)

Tool for comparing two `CSV` files, and create the output of the difference. Ngos can compare millions of lines of csv file just in second. 

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
$ ngos -old old_file.csv -new new_file.csv -o output.csv
```

Which is

- `old_file.csv`

```csv
Los Angeles,34°03′N,118°15′W
New York City,40°42′46″N,74°00′21″W
Paris,48°51′24″N,2°21′03″E
```

- `new_file.csv`

```csv
Los Angeles,34°03′N,118°15′W
New York City,40°42′46″N,74°00′21″W
Paris,48°51′24″N,2°21′03″E
Jakarta,48°51′24″N,2°21′03″W
```

#
- `output.csv`

```csv
Jakarta,48°51′24″N,2°21′03″W
```
