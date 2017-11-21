package faker

import (
	"github.com/Gujarats/faker/provider/en_US"
)

var genders = []string{"male", "female"}

func Names(n int) []string {
	var names []string
	for i := 0; i < n; i++ {
		names = append(names, Name())
	}

	return names
}

// Get a random name
func Name() string {
	PersonGenerator := []string{PersonSuffix(), Person()}
	return PersonGenerator[getRandomIndex(PersonGenerator)]
}

// generate person with kind of rand suffix,male ,female
func PersonSuffix() string {
	gender := genders[getRandomIndex(genders)]
	if gender == "male" {
		return PersonMale()
	} else {
		return PersonFemale()
	}
}

// generate person without suffix
func Person() string {
	gender := genders[getRandomIndex(genders)]
	if gender == "male" {
		return Male()
	} else {
		return Female()
	}
}

func PersonMale() string {
	return Male() + " " + suffix()
}

func PersonFemale() string {
	return Female() + " " + suffix()
}

// return random male name with first and last name
func Male() string {
	return firstNameMale() + " " + lastName()
}

func firstNameMale() string {
	return getRandomValue(en_US.FirstNameMale)
}

func Female() string {
	return firstNameFemale() + " " + lastName()
}

func firstNameFemale() string {
	return getRandomValue(en_US.FirstNameFemale)
}

func lastName() string {
	return getRandomValue(en_US.LastName)
}

func suffix() string {
	return getRandomValue(en_US.Suffix)
}
