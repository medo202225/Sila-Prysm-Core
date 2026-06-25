package testutil

import "github.com/sila-chain/Sila-Prysm-Core/v7/encoding/ssz/query"

type PathTest struct {
	Path     string
	Expected any
}

type TestSpec struct {
	Name      string
	Type      query.SSZObject
	Instance  query.SSZObject
	PathTests []PathTest
}
