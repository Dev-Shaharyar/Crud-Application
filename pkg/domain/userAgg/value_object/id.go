package value_object

// ID represents a user identifier
type ID string

func (id ID) String() string {
	return string(id)
}
