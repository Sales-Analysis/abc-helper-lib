package abc

// Private type for sorting
type byValue []pair

// Sorting helpers (Len, Swap, Less)
func (a byValue) Len() int           { return len(a) }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValue) Less(i, j int) bool { return a[i].value > a[j].value }
