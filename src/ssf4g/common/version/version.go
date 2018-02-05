package version

import (
	"github.com/hashicorp/go-version"
)

func VersionLessThan(vera, verb string) (bool, error) {
	versionA, err := version.NewVersion(vera)

	if err != nil {
		return false, err
	}

	versionB, err := version.NewVersion(verb)

	if err != nil {
		return false, err
	}

	if versionA.LessThan(versionB) == true {
		return true, nil
	}

	return false, nil
}

func VersionGreaterThan(vera, verb string) (bool, error) {
	versionA, err := version.NewVersion(vera)

	if err != nil {
		return false, err
	}

	versionB, err := version.NewVersion(verb)

	if err != nil {
		return false, err
	}

	if versionA.GreaterThan(versionB) == true {
		return true, nil
	}

	return false, nil
}
