package allergies

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	lst := []string{}
	for k, v := range allergens {
		if allergies&v == v {
			lst = append(lst, k)
		}
	}
	return lst
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&allergens[allergen] == allergens[allergen]
}
