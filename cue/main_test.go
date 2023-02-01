package main

import (
	"errors"
	"os"
	"strings"
	"testing"

	"cuelang.org/go/cue/cuecontext"
	cerrors "cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
	cjson "cuelang.org/go/encoding/json"
)

func TestValidateOrganisation(t *testing.T) {
	// If there are multiple invalid fields only the first one will be reported.
	// TODO: Check if there is a workaround.

	tests := map[string]struct {
		filename    string
		expErr      bool
		invalidPath string
	}{
		"ValidOrg": {
			filename: "testdata/org.json",
		},
		"InvalidUUIDField": {
			filename:    "testdata/org-invalid-uuid.json",
			expErr:      true,
			invalidPath: "#Organisation.uuid",
		},
		"InvalidYearFoundedField": {
			filename:    "testdata/org-invalid-int.json",
			expErr:      true,
			invalidPath: "#Organisation.yearFounded",
		},
	}

	ctx := cuecontext.New()

	binst := load.Instances([]string{"concepts-schema.cue"}, nil)
	vals, err := ctx.BuildInstances(binst)
	if err != nil {
		t.Fatalf("loading failed: %s", err)
	}

	orgsSchema := vals[0].LookupDef("#Organisation")

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			data := readFileHelper(t, test.filename)
			err = cjson.Validate(data, orgsSchema)

			if err != nil && !test.expErr {
				t.Fatalf("did not expect error, got: %s", err)
			}
			if err == nil && test.expErr {
				t.Fatalf("expected error, did not get one")
			}

			if err == nil {
				return
			}

			var cerr cerrors.Error
			errors.As(err, &cerr)

			got := strings.Join(cerr.Path(), ".")
			if got != test.invalidPath {
				t.Fatalf("got: %s, want: %s", got, test.invalidPath)
			}
		})
	}
}

func readFileHelper(t *testing.T, filename string) []byte {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("reading %s: %s", filename, err)
	}

	return data
}
