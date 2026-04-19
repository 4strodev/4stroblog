package models

// SiteMeta hold information about the site itself
// like if it's first start or not
// with this we can make some tasks like create admin user
type SiteMeta struct {
	FirstStart bool
}
