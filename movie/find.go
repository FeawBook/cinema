package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "1234":
		return "Arita"
	case "3344":
		return "Ongbak"
	}
	return "not found"
}
