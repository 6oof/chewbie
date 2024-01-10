package helpers

// Seo represents the SEO (Search Engine Optimization) metadata for a web page.
type Seo struct {
	// Name is the name of the web page.
	Name string

	// Title is the title of the web page.
	Title string

	// Description is the meta description for the web page.
	Description string

	// Keywords are the meta keywords for the web page.
	Keywords string
}

// BaseSeo returns a default Seo object with values retrieved from environmental variables.
//
// The default values are used if the corresponding environmental variable is not set.
// - NAME: The name of the web page. Default is "Miniweb".
// - DESCRIPTION: The meta description for the web page. Default is "Miniweb".
//
// Example Usage:
//
//	seo := BaseSeo()
//	fmt.Println(seo.Name) // Output: Miniweb
//	fmt.Println(seo.Description) // Output: Miniweb
func BaseSeo() Seo {
	return Seo{
		Name:        Env("NAME", "Miniweb"),
		Description: Env("DESCRIPTION", "Miniweb"),
		// Title and Keywords can be added similarly using Env() for each.
	}
}
