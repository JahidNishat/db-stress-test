package main

type User struct {
	Name string
}

func stayOnStack() User {
	u := User{Name: "Jahid"}
	return u
}

func escapeToHeap() *User {
	u := &User{Name: "Jahid"}
	return u
}

func escape() {
	_ = stayOnStack()
	_ = escapeToHeap()
}
