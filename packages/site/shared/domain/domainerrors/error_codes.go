package domainerrors

type errorCode int

const (
	// Used when an entity should be found and was not
	ENTITY_NOT_FOUND errorCode = iota
	// Used when due domain business rules creating or removing data
	// is not possible
	DATA_CONFLICT
	// Used when a context was finished before operation succeed
	CONTEXT_FINISHED
	// Got while executing some functionality inside the self application
	// that returned an unexpected error. Differs from DATABASE because
	// the error occurred inside the app not on external components
	RUNTIME
	// Used when got an unexpected database error
	DATABASE
)
