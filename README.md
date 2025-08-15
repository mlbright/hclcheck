# hclcheck

`hclcheck` is a tool that validates .tfvars files in which all lines are initially commented out.

## Usage

```bash
hclcheck <file>
```

### Example

```bash
find ~/Source/ccloud-cli/main -type f -name "*.tfvars.example" | xargs -I % hclcheck %
```

## Building

To build the `hclcheck` binary, you need to have Go installed. Then, run the following command in the root directory of the project:

```bash
go build -o hclcheck main.go
```

Update dependencies with:

```bash
go get -u ./...
```


```bash
go mod tidy
```
