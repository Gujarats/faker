package faker

import (
	"github.com/Gujarats/faker/provider/en_US"
)

var gender = []string{"male", "female"}

// Get a random name for specific male or female if specify
func Name() string {
	//genderType := getRandomIndex(gender)
	return ""
}

func Person() string {
	// generate person with kind of rand prefix male and female and degree
	_ = Prefix()

	return ""
}

func Prefix() string {
	return ""

}

func Degree() string {
	return ""

}

// return random male name with first and last name
func male() string {
	firstNameMale := en_US.FirstNameMale[getRandomIndex(en_US.FirstNameMale)]
	lastNameMale := en_US.LastName[getRandomIndex(en_US.LastName)]

	return firstNameMale + " " + lastNameMale
}
