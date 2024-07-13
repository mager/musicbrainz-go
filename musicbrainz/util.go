package musicbrainz

func IncludesContains(inc Includes, i Include) bool {
	for _, st := range inc {
		if st == i {
			return true
		}
	}
	return false
}
