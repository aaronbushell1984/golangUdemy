package marshal

import "fmt"

func ExampleMakePerson() {
	jb := MakePerson("James", "Bond", 32)
	mp := MakePerson("Miss", "Moneypenny", 25)
	fmt.Println(jb)
	fmt.Println(mp)
	// Output:
	// {James Bond 32}
	// {Miss Moneypenny 25}
}

func ExampleAddPerson() {
	jb := MakePerson("James", "Bond", 32)
	mp := MakePerson("Miss", "Moneypenny", 25)
	people := []Person{}
	fmt.Println(AddPerson(people, jb, mp))
	// Output:
	// [{James Bond 32} {Miss Moneypenny 25}]
}

func ExampleMarshalAsString() {
	jb := MakePerson("James", "Bond", 32)
	mp := MakePerson("Miss", "Moneypenny", 25)
	people := []Person{}
	p := AddPerson(people, jb, mp)
	fmt.Println(MarshalAsString(p))
	// Output:
	// [{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":25}] <nil>
}

func ExampleMarshalAsByte() {
	jb := MakePerson("James", "Bond", 32)
	mp := MakePerson("Miss", "Moneypenny", 25)
	people := []Person{}
	p := AddPerson(people, jb, mp)
	fmt.Println(MarshalAsByte(p))
	// Output:
	// [91 123 34 70 105 114 115 116 34 58 34 74 97 109 101 115 34 44 34 76 97 115 116 34 58 34 66 111 110 100 34 44 34 65 103 101 34 58 51 50 125 44 123 34 70 105 114 115 116 34 58 34 77 105 115 115 34 44 34 76 97 115 116 34 58 34 77 111 110 101 121 112 101 110 110 121 34 44 34 65 103 101 34 58 50 53 125 93] <nil>
}
