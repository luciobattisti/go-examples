# Golang Basic Project with local module import

# Set workspaces
```bash
cd ../
go work init
go work use hello
```

# Init main module
```bash
cd hello
go mod init hello
```

# Init helper module
```bash
cd helper
go mod init examples.com/helper
```

# Rename helper module for local access
```bash
cd ../
go mod edit -replace examples.com/helper=./helper
go get examples.com/helper
```

# Build
```bash
go build -o hello.o
```