package main

var (
	categories = []string{
		"osint",
		"exploit",
		"post-exploit",
	}

	register = []struct {
		name        string
		category    string
		description string
		path        string
	}{
		{
			name:        "DetectDee",
			category:    "osint",
			description: "Hunt down social media accounts by username, email or phone across social networks.",
			path:        "detectdee/detectdee.go",
		},
		{
			name:        "DetectDee",
			category:    "osint",
			description: "Hunt down social media accounts by username, email or phone across social networks.",
			path:        "detectdee/detectdee.go",
		},
	}
)

func listAllModules() []struct {
	name        string
	category    string
	description string
	path        string
} {
	return register
}

func listCategory(cat string) []struct {
	name        string
	category    string
	description string
	path        string
} {
	validCategory := false

	for _, v := range categories {
		if v == cat {
			validCategory = true
		}
	}

	modules := []struct {
		name        string
		category    string
		description string
		path        string
	}{}

	if !validCategory {
		return modules
	}

	for _, v := range register {
		if v.category == cat {
			modules = append(modules, v)
		}
	}

	return modules

}
