package unmarshal

import (
	"fmt"
	"github.com/aaronbushell1984/golangUdemy/18/json/marshal"
)

func ExampleMakePerson() {
	jb := marshal.MakePerson("James", "Bond", 32)
	mp := marshal.MakePerson("Miss", "Moneypenny", 25)
	fmt.Println(jb)
	fmt.Println(mp)
	// Output:
	// {James Bond 32}
	// {Miss Moneypenny 25}
}

func ExampleAddPerson() {
	jb := marshal.MakePerson("James", "Bond", 32)
	mp := marshal.MakePerson("Miss", "Moneypenny", 25)
	people := []marshal.Person{}
	fmt.Println(marshal.AddPerson(people, jb, mp))
	// Output:
	// [{James Bond 32} {Miss Moneypenny 25}]
}

func ExampleUnmarshal() {
	dataAllFields := []byte(`[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":25}]`)
	var pplAllFields people
	dataNoAge := []byte(`[{"First":"James","Last":"Bond"},{"First":"Miss","Last":"Moneypenny"}]`)
	var pplNoAge people
	fmt.Println(Unmarshal(dataAllFields, pplAllFields))
	fmt.Println(Unmarshal(dataNoAge, pplNoAge))
	// Output:
	// [{James Bond 32 } {Miss Moneypenny 25 }] <nil>
	// [{James Bond 0 } {Miss Moneypenny 0 }] <nil>
}