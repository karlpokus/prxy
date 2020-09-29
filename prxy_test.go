package prxy

import (
	"testing"
	"strings"

	"github.com/karlpokus/bufw"
)

type test struct {
	Name, Src, Dest string
	Args []string
	Err bool
}

func TestValidation(t *testing.T) {
	testTable := []test{
		{
			Name: "wrong len",
			Args: []string{"foo"},
			Src: "",
			Dest: "",
			Err: true,
		},
		{
			Name: "empty src",
			Args: []string{"", "doo"},
			Src: "",
			Dest: "",
			Err: true,
		},
		{
			Name: "empty dest",
			Args: []string{"foo", ""},
			Src: "",
			Dest: "",
			Err: true,
		},
		{
			Name: "wrong format",
			Args: []string{"foo:1", "doo"},
			Src: "",
			Dest: "",
			Err: true,
		},
		{
			Name: "correct format",
			Args: []string{"foo:1", "doo:2"},
			Src: "foo:1",
			Dest: "doo:2",
			Err: false,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T){
			src, dest, err := validateArgs(tt.Args)
			if src != tt.Src {
				t.Fatalf("expected %s and %s to match", src, tt.Src)
			}
			if dest != tt.Dest {
				t.Fatalf("expected %s and %s to match", dest, tt.Dest)
			}
			if tt.Err && err != argErr {
				t.Fatal("expect errors to match")
			}
			if !tt.Err && err != nil {
				t.Fatal("expect errors to match")
			}
		})
	}
}

// client writes to server
func TestCopy(t *testing.T) {
	client := bufw.New()
	in := "something"
	server := strings.NewReader(in)
	go copy(client, server, "", "")
	client.Wait()
	out := client.String()
	if in != out {
		t.Fatalf("expected %s and %s to match", in, out)
	}
}
