package Resources
type Resources struct {
	Name string
	Verbs []string
}
type GroupResources struct {
	Group string
	Version string
	Resources []*Resources
}