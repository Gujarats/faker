package faker

import (
	"reflect"
	"runtime"
	"testing"
)

// get random 10 male's name
func TestMale(t *testing.T) {
	names := make(map[string]bool)
	for i := 0; i < 10; i++ {
		name := Male()
		if _, ok := names[name]; ok {
			t.Errorf("the males() function return the duplicate value or previous generated value ")
		} else {
			names[name] = true
		}
	}

}

func TestName(t *testing.T) {
	names := make(map[string]bool)
	for i := 0; i < 10; i++ {
		name := Name()
		if _, ok := names[name]; ok {
			t.Errorf("the males() function return the duplicate value or previous generated value ")
		} else {
			names[name] = true
		}
	}
}

func TestNames(t *testing.T) {
	totals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 30, 100}
	for index, total := range totals {
		names := Names(total)
		if hasDuplicate(names) {
			t.Errorf("index = %v,Generator Names(%+v) return duplicate data = %+v\n", index, total, names)
		}
	}
}

func TestEmptyValueForFunc(t *testing.T) {
	var testFunc = []func() string{Name, PersonSuffix, Person, PersonMale, PersonFemale, suffix, Male, Female, lastName}
	for i := 0; i < len(testFunc); i++ {
		result := testFunc[i]()
		if result == "" || len(result) <= 0 {
			funcName := runtime.FuncForPC(reflect.ValueOf(testFunc[i]).Pointer()).Name()
			t.Errorf("%+s function return empty", funcName)
		}
	}
}

func TestCheckHasDuplicate(t *testing.T) {
	testObjects := []struct {
		datas    []string
		expected bool
	}{
		{
			datas:    []string{"a", "b", "c", "d", "2", "3"},
			expected: false,
		},
		{
			datas:    []string{"a", "b", "c", "a", "2", "3"},
			expected: true,
		},
		{
			datas:    []string{"aa", "b", "c", "aa", "2", "3"},
			expected: true,
		},
		{
			datas:    []string{"aa", "aa", "bb"},
			expected: true,
		},
	}

	for index, testObject := range testObjects {
		result := hasDuplicate(testObject.datas)
		if result != testObject.expected {
			t.Errorf("index = %v, expected = %+v, actual = %+v\n", index, testObject.expected, result)
		}
	}
}

func hasDuplicate(datas []string) bool {
	storage := make(map[string]bool)
	for _, data := range datas {
		if _, ok := storage[data]; ok {
			return true
		} else {
			storage[data] = true
		}

	}

	return false
}
