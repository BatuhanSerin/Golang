## Movie Library

This project made with Golang and contains a list of movies and movies can be searched.

### Usage

##### 1 - To show the list of movies
```go
go run main.go list
```
##### 2 - To search the movie in a list
```go
go run main.go search <bookName>
```
The search command is not case sensitive.
###### Example
```go
go run main.go search Requiem For A Dream 
go run main.go search requiem for a dream
```

