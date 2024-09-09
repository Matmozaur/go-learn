package main

import (
	"fmt"
	"os"
	"strings"
)

func SplitExt(path string) (string, string) {
	i := strings.LastIndex(path, ".")
	if i == -1 {
		return path, ""
	}

	return path[:i], path[i:]
}

func Sum(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func Mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Mean of empty slice")
	}

	return Sum(values) / float64(len(values)), nil
}

func fileHead(fileName string, size int) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, size)
	n, err := file.Read(buf)

	if err != nil {
		return nil, err
	}
	if n != size {
		return nil, fmt.Errorf("%q too small", fileName)
	}

	return buf, nil
}

type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func Promote(u User, r Role) {
	u.Role = r
}

func myDiag(arr [][]int) (int, error) {
	n, m := len(arr), len(arr[0])
	res := 1
	for k := 0; k < min(n, m); k++ {
		res *= arr[k][k]
	}
	// for k := range min(n, m) {
	// 	res *= arr[k][k]
	// }
	return res, nil
}

func main() {
	root, ext := SplitExt("app.go")
	fmt.Printf("root=%#v, path=%#v\n", root, ext)

	values := []float64{2, 4, 8}
	m, err := Mean(values)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(m)

	data, err := fileHead("head.png", 8)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(data)
	}

	u := User{"elliot", Viewer}
	Promote(u, Admin)
	fmt.Printf("%#v\n", u)

	a := Role("a")
	fmt.Printf("%#v\n", a)

	r, _ := myDiag([][]int{{2, 1, 3}, {2, 2, 2}, {3, 3, 3}})
	fmt.Printf("%#v\n", r)
}
