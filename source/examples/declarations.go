package main

type ExampleTypeString string

type ExampleTypeStruct struct {
	val     string
	version int
}

type ExampleTypeInterface interface {
	Value() string
}

type ExampleTypeAlias = ExampleTypeTwo

type (
	ExampleGroupedTypeString string

	ExampleGroupedTypeStruct struct {
		val     string
		version int
	}

	ExampleGroupedTypeAliasDeclaration = ExampleGroupedTypeStruct

	ExampleGroupedTypeInterface interface {
		Value() string
	}
)
