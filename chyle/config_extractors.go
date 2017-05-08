package chyle

import (
	"github.com/antham/envh"
)

// extractorsConfigurator validates jira config
// defined through environment variables
type extractorsConfigurator struct {
	chyleConfig *CHYLE
	config      *envh.EnvTree
	definedKeys []string
}

func (e *extractorsConfigurator) process() (bool, error) {
	if e.isDisabled() {
		return true, nil
	}

	for _, f := range []func() error{
		e.validateExtractors,
	} {
		if err := f(); err != nil {
			return true, err
		}
	}

	e.setExtractors()

	return true, nil
}

// isDisabled checks if matchers are enabled
func (e *extractorsConfigurator) isDisabled() bool {
	return featureDisabled(e.config, [][]string{{"CHYLE", "EXTRACTORS"}})
}

// validateExtractors checks threesome extractor fields
func (e *extractorsConfigurator) validateExtractors() error {
	for _, key := range e.config.FindChildrenKeysUnsecured("CHYLE", "EXTRACTORS") {
		if err := validateSubConfigPool(e.config, []string{"CHYLE", "EXTRACTORS", key}, []string{"ORIGKEY", "DESTKEY", "REG"}); err != nil {
			return err
		}

		if err := validateRegexp(e.config, []string{"CHYLE", "EXTRACTORS", key, "REG"}); err != nil {
			return err
		}

		e.definedKeys = append(e.definedKeys, key)
	}

	return nil
}

// setExtractors update chyleConfig with extracted extractors
func (e *extractorsConfigurator) setExtractors() {
	e.chyleConfig.EXTRACTORS = map[string]map[string]string{}

	for _, key := range e.definedKeys {
		e.chyleConfig.EXTRACTORS[key] = map[string]string{}

		for _, field := range []string{"ORIGKEY", "DESTKEY", "REG"} {
			e.chyleConfig.EXTRACTORS[key][field] = e.config.FindStringUnsecured("CHYLE", "EXTRACTORS", key, field)
		}
	}
}