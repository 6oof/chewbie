package helpers

type Seo struct {
	Name        string
	Title       string
	Description string
	Keywords    string
}

func BaseSeo() Seo {
	return Seo{
		Name:        Env("NAME", "Miniweb"),
		Description: Env("DESCRIPTION", "Miniweb"),
	}
}
