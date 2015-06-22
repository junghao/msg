package msg

import (
	"fmt"
)

// uaTags generates tags for sending to Urban Airship.
func (v *VAL) uaTags() (tags []string) {
	if v.err != nil {
		return
	}

	// generate all levels above alert for given volcano, and all
	for i := v.VolcanicAlert.Level; i < len(VolcanicAlertLevels); i++ {
		tags = append(tags, fmt.Sprintf("%d@%s", i, v.Volcano.VolcanoID))
		tags = append(tags, fmt.Sprintf("%d@all", i))
	}

	return
}
