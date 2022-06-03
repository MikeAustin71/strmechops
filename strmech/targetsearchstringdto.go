package strmech

// TargetSearchStringDto - A Target Search String Data Transfer
// Object. This type is designed to transfer a rune array. This
// rune array can be used to carry out text character searches.
type TargetSearchStringDto struct {
	CharsToSearch []rune
}
