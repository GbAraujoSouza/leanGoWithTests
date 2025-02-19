package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	City    string
	Age     int
	Profile Profile
}

type Profile struct {
	Email  string
	Rating float64
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{Name: "Gabriel"},
			[]string{"Gabriel"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Gabriel", "London"},
			[]string{"Gabriel", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				City string
				Age  int
			}{"Gabriel", "London", 24},
			[]string{"Gabriel", "London"},
		},
		{
			"struct with nested fields",
			Person{
				Name: "Gabriel",
				City: "London",
				Age:  24,
				Profile: Profile{
					"email",
					4.9,
				}},
			[]string{"Gabriel", "London", "email"},
		},
		{
			"pointers to things",
			&Person{
				Name: "Gabriel",
				City: "London",
				Age:  24,
				Profile: Profile{
					"email",
					4.9,
				}},
			[]string{"Gabriel", "London", "email"},
		},
		{
			"struct with slices",
			[]Person{
				{
					Name: "Gabriel",
					City: "London",
					Age:  24,
					Profile: Profile{
						"email",
						4.9,
					},
				},
				{
					Name: "Larissa",
					City: "Paris",
					Age:  25,
					Profile: Profile{
						"email2",
						5.0,
					},
				},
			},
			[]string{"Gabriel", "London", "email", "Larissa", "Paris", "email2"},
		},
		{
			"arrays",
			[2]Profile{
				{
					"email1",
					4.9,
				},
				{
					"email2",
					4.9,
				},
			},
			[]string{"email1", "email2"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"bolo":  "de morango",
			"torta": "de limao",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

        assertContains(t, got, "de morango")
        assertContains(t, got, "de limao")
	})

    t.Run("with channels", func(t *testing.T) {
        aChannel := make(chan Profile)

        go func() {
            aChannel <- Profile{"email1", 2.8}     
            aChannel <- Profile{"email2", 5.0}     
            close(aChannel)
        }()

        var got []string
        want := []string{"email1", "email2"}

        walk(aChannel, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("with functions", func(t *testing.T) {
        aFunction := func () (Profile, Profile) {
            return Profile{"email1", 2.8}, Profile{"email2", 5.0} 
        }

        var got []string
        want := []string{"email1", "email2"}

        walk(aFunction, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	// linear search
	contains := false
	for _, s := range haystack {
		if s == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}

}

