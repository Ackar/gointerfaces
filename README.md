# gointerfaces

`gointerfaces` is a simple binary to list exported interfaces from a package.

### Installation

```
go get -u github.com/Ackar/gointerfaces
```

### Example usage

```
$ gointerfaces $(go env GOROOT)/src/encoding/
encoding.BinaryMarshaler
encoding.BinaryUnmarshaler
encoding.TextMarshaler
encoding.TextUnmarshaler
```