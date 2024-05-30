package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	// score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string
	for _, char := range s.name {
		char += 3 
		if char > 122 {
			char -= 26
		}
		nameEncode += string(char)

	}
	s.nameEncode = nameEncode
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode string
	for _, char := range s.nameEncode {
		char -= 3
		if char < 97 {
			char += 26
		}
		nameDecode += string(char)
	}
	return nameDecode
}


func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Println("[1] Encrypt \n[2] Decrypt \nChoose your menu")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nEncode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Print("Wrong Input Name Menu")
	}

}