package set

type Set interface {
	Equal(val1, val2 interface{}) bool
	Elements() []interface{}
	Add(val interface{}) interface{}
	Len() int
}

func Add(set Set, val interface{}) interface{} {
	for _, s := range set.Elements() {
		if set.Equal(s, val) {
			return set
		}
	}
	return set.Add(val)
}

// todo. This Method is not complete.
func Remove(set Set, val interface{}) interface{} {
	for _, s := range set.Elements() {
		if set.Equal(s, val) {
			break
		}
	}
	return nil
}

func Len(set Set) int {
	return set.Len()
}

type StringSlice []string

func (s StringSlice) Equal(val1, val2 interface{}) bool { return val1.(string) == val2.(string) }
func (s StringSlice) Elements() []interface{} {
	var strs = []string(s)
	var result = make([]interface{}, len(strs))
	for _, s := range strs {
		result = append(result, s)
	}
	return result
}
func (s StringSlice) Add(val interface{}) interface{} {
	s = append([]string(s), val.(string))
	return s
}
func (s StringSlice) Len() int { return len([]string(s)) }

func AddToStrings(set []string, val string) []string {
	return []string((Add(StringSlice(set), val).(StringSlice)))
}

func RemoveFromStrings(set []string, val string) {
	Remove(StringSlice(set), val)
}
