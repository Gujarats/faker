package faker

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// get a random value with given format
func Namesf(n int, formats ...string) []string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, Namef(formats...))
	}

	return result
}

// get a random value with given format
func Namef(formats ...string) string {
	if len(formats) == 0 {
		log.Fatalf("format = %+v must not empty", formats)
	}

	validFormats, err := getValidFormat(formats...)
	if err != nil {
		log.Fatal(err)
	}

	formatter := getFormatter()
	var result string
	for _, validFormat := range validFormats {
		result = result + " " + formatter[validFormat]()
	}

	return result
}

func getValidFormat(formats ...string) ([]string, error) {
	if len(formats) == 0 {
		message := fmt.Sprintf("format = %+v must not empty", formats)
		return nil, errors.New(message)
	}

	formats = lowerCaseFormats(formats...)

	formatter := getFormatter()
	for _, format := range formats {
		if _, ok := formatter[format]; !ok {
			message := fmt.Sprintf("undefined format = %+v\n use this one of theses formats instead = %+v\n", format, getFormatList())
			return nil, errors.New(message)
		}
	}

	return formats, nil
}

func getFormatList() []string {
	formatter := getFormatter()
	var formatList []string
	for key, _ := range formatter {
		formatList = append(formatList, key)
	}
	return formatList
}

// Define all format that we can use and speciy to specifi function
func getFormatter() map[string]func() string {
	formatter := make(map[string]func() string)
	formatter["male"] = Male
	formatter["female"] = Female
	formatter["suffix"] = suffix
	formatter["first-name-female"] = firstNameFemale
	formatter["first-name-male"] = firstNameMale
	formatter["last-name"] = lastName
	return formatter
}

func lowerCaseFormats(formats ...string) []string {
	var result []string
	if len(formats) > 0 {
		for _, format := range formats {
			result = append(result, strings.ToLower(format))
		}
	}

	return result
}
