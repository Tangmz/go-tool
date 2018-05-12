package set

//	The Set was implemented by golang.
type Set interface {
	// Determine if two values are equal. if it's true, val1 is equal to val2.
	Equal(val1, val2 interface{}) bool

	// Return the []interface{} that each element is the element of the origin Set.
	Elements() []interface{}

	// Add an unique element to Set.
	Add(val interface{}) interface{}

	// Remove the element from Set.
	Remove(val interface{}) interface{}

	// Return the length of the Set.
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

func Remove(set Set, val interface{}) interface{} {
	return set.Remove(val)
}

func Len(set Set) int {
	return set.Len()
}

type StringSlice []string

func (ss StringSlice) Equal(val1, val2 interface{}) bool { return val1.(string) == val2.(string) }
func (ss StringSlice) Elements() []interface{} {
	var strs = []string(ss)
	var result = []interface{}{}
	for _, s := range strs {
		result = append(result, s)
	}
	return result
}
func (ss StringSlice) Add(val interface{}) interface{} {
	var result = append([]string(ss), val.(string))
	return StringSlice(result)
}
func (ss StringSlice) Len() int { return len([]string(ss)) }
func (ss StringSlice) Remove(val interface{}) interface{} {
	var strs = []string(ss)
	var result = []string{}
	for _, s := range strs {
		if ss.Equal(s, val) {
			continue
		}
		result = append(result, s)
	}
	return result
}

func AddToStrings(set []string, val string) []string {
	var result = Add(StringSlice(set), val).(StringSlice)
	return []string(result)
}

func RemoveFromStrings(set []string, val string) []string {
	return Remove(StringSlice(set), val).([]string)
}
