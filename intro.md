# Golang Project

- **creating project** : `go mod init <project_name>`
- `go build -o bin/main main.go`: build object file and store in bin folder
- `go mod tidy`: it will find all the required packages and download them in go module cache directory
- Makefile is used to automate script such as build, run etc and we use `make <task_name>` to run the script

## Use of pointer in golang
**Use Pointers (\*T) when**:

- Mutating data: The function needs to change the original struct.

- Avoiding large copies: The struct is large, and passing a pointer is faster.

- Handling optional fields: You need a nil state to distinguish between zero and unset values.

**Use Values (T) when**:

Working with built-in references: Maps, slices, and channels are already references—never point to them (*[]T).

Using small structs: Simple types (like a 2D point) are faster to copy than allocate on the heap.

Enforcing immutability: You want a safe local copy to prevent unintended side effects