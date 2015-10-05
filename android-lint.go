package pipeline

import "fmt"

type AndroidLint struct {
	LintFile            string
	HasAdvancedSettings bool
	FailedTotalNormal   float64
	FailedTotalHigh     float64
}

func (androidLint *AndroidLint) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "lintFile":
			androidLint.LintFile = value.(string)
		case "failedTotalNormal":
			androidLint.HasAdvancedSettings = true
			androidLint.FailedTotalNormal = value.(float64)
		case "failedTotalHigh":
			androidLint.HasAdvancedSettings = true
			androidLint.FailedTotalHigh = value.(float64)
		default:
			return fmt.Errorf("Unknown key for AndroidLint plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}
