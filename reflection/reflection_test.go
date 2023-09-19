package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Sam"},
			[]string{"Sam"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Sam", "Doncaster"},
			[]string{"Sam", "Doncaster"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sam", 30},
			[]string{"Sam"},
		},
		{
			"struct with nested fields",
			Person{
				"Sam",
				Profile{
					Age:  30,
					City: "Doncaster",
				},
			},
			[]string{"Sam", "Doncaster"},
		},
		{
			"struct with pointer to things",
			&Person{
				"Sam",
				Profile{30, "Doncaster"},
			},
			[]string{"Sam", "Doncaster"},
		},
		{
			"slices of structs",
			[]Profile{
				{30, "Doncaster"},
				{33, "London"},
			},
			[]string{"Doncaster", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
