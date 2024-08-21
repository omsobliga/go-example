## cmd

step1
```
protoc --go_out=. addressbook.proto
```

step2
```
go run cmd/add_people/add_people.go {filename}
```

step3
```
go run cmd/list_people/list_people.go {filename}
```
