package npm

import (
	"os"
	"path"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9133/go-dep-parser/pkg/types"
)

func TestParse(t *testing.T) {
	vectors := []struct {
		file string // Test input file
		want []types.Library
	}{
		{
			file: "testdata/package-lock_normal.json",
			want: NpmNormal,
		},
		{
			file: "testdata/package-lock_react.json",
			want: NpmReact,
		},
		{
			file: "testdata/package-lock_with_dev.json",
			want: NpmWithDev,
		},
		{
			file: "testdata/package-lock_many.json",
			want: NpmMany,
		},
		{
			file: "testdata/package-lock_nested.json",
			want: NpmNested,
		},
	}

	for _, v := range vectors {
		t.Run(path.Base(v.file), func(t *testing.T) {
			f, err := os.Open(v.file)
			require.NoError(t, err)

			got, err := Parse(f)
			require.NoError(t, err)

			sortLibs(got)
			sortLibs(v.want)

			assert.Equal(t, v.want, got)
		})
	}
}

func sortLibs(libs []types.Library) {
	sort.Slice(libs, func(i, j int) bool {
		ret := strings.Compare(libs[i].Name, libs[j].Name)
		if ret == 0 {
			return libs[i].Version < libs[j].Version
		}
		return ret < 0
	})
}
