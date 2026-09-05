package common

// GetString safely dereferences a string pointer, returning an empty string if the pointer is nil.
func GetString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
