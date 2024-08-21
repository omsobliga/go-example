package handler

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Book struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
