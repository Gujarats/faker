package faker

import (
	"github.com/Gujarats/faker/provider/en_US"
)

var genders = []string{"male", "female"}

// Get a random name for specific male or female if specify
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

func suffix() string {
	return en_US.Suffix[getRandomIndex(en_US.Suffix)]
}

// return random male name with first and last name
func Male() string {
	firstNameMale := en_US.FirstNameMale[getRandomIndex(en_US.FirstNameMale)]
	return firstNameMale + " " + lastName()
}

func Female() string {
	firstNameFemale := en_US.FirstNameFemale[getRandomIndex(en_US.FirstNameFemale)]
	return firstNameFemale + " " + lastName()
}

func lastName() string {
	return en_US.LastName[getRandomIndex(en_US.LastName)]
}
