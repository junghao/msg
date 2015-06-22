package msg

import (
	"fmt"
	"sort"
	"testing"
)

func TestUAVALTags(t *testing.T) {
	expectedTags := []string{
		`2@all`,
		`3@all`,
		`4@all`,
		`5@all`,
		`2@mayorisland`,
		`3@mayorisland`,
		`4@mayorisland`,
		`5@mayorisland`,
	}

	v := &VAL{VolcanicAlert: VolcanicAlert{Level: 2, Activity: `test`, Hazards: `test`}, Volcano: Volcano{VolcanoID: `mayorisland`, Title: `test`}, err: nil}
	tags := v.uaTags()

	sort.Strings(tags)
	sort.Strings(expectedTags)

	if fmt.Sprintf("%v", expectedTags) != fmt.Sprintf("%v", tags) {
		t.Error("Didn't get expected tags")
	}

}
