# LINT
```
go get -u github.com/golang/lint/golint
```

# Tools
```
go get -u golang.org/x/tools
```

# CTags
```
go get -u github.com/jstemmer/gotags
```

And then generate tags simply by invoking it like this from the root of your source code.
```
gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .
```

# Data Types
```
bool
int int8 int16 int32 int64
uint unint8 unint16 unint32 unint64

byte // alias for int8
rune  // alias for int32
      // represents an unicode code point

float32 float64
complex64 complex128
```
