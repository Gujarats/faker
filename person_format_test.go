package faker

import "testing"

func TestGetPersonFormat(t *testing.T) {
	testObjects := []struct {
		formats []string
	}{
		{
			formats: []string{"male", "suffix"},
		},

		{
			formats: []string{"female", "suffix"},
		},
		{
			formats: []string{"female", "last-name"},
		},
		{
			formats: []string{"male", "last-name"},
		},
		{
			formats: []string{"first-name-male", "last-name"},
		},
		{
			formats: []string{"first-name-female", "last-name"},
		},
	}

	for index, testObject := range testObjects {
		result := GetPersonFormat(testObject.formats...)
		if len(result) == 0 {
			t.Errorf("index = %v, result = %+v must not empty", index, result)
		}
	}
}

func TestIsValidFormat(t *testing.T) {
	testObjects := []struct {
		formats []string
		ok      bool
	}{
		{
			formats: []string{"male", "suffix"},
			ok:      true,
		},
		{
			formats: []string{"malemale", "suffix"},
			ok:      false,
		},
		{
			formats: []string{"Male", "suffix"},
			ok:      true,
		},
		{
			formats: []string{"MalE", "suFfix"},
			ok:      true,
		},
		{
			formats: []string{"first-Name-MalE", "suFfix"},
			ok:      true,
		},
	}

	for index, testObject := range testObjects {
		_, err := getValidFormat(testObject.formats...)
		if testObject.ok {
			if err != nil {
				t.Errorf("index = %v, expected= %+v ,actual = %+v\n", index, nil, err)
			}
		}
	}
}
