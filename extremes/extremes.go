package extremes

func Min(a, b int) int {

	switch {
	case a < b:
		return a
	default:
		return b
	}
}

func Max(a, b int) int {

	switch {
	case a > b:
		return a
	default:
		return b
	}
}

func Min64(a, b int64) int64 {

	switch {
	case a < b:
		return a
	default:
		return b
	}
}

func Max64(a, b int64) int64 {

	switch {
	case a > b:
		return a
	default:
		return b
	}
}
