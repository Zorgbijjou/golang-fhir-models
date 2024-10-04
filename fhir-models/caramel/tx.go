package caramel

import (
	"encoding/json"
	"github.com/zorgbijjou/golang-fhir-models/fhir-models/caramel/to"
	"github.com/zorgbijjou/golang-fhir-models/fhir-models/fhir"
)

type TransactionBuilder fhir.Bundle

func Transaction() *TransactionBuilder {
	return &TransactionBuilder{
		Type: fhir.BundleTypeTransaction,
	}
}

func (t *TransactionBuilder) Create(resource interface{}, opts ...BundleEntryOption) *TransactionBuilder {
	return t.addEntry(resource, ResourceType(resource), fhir.HTTPVerbPOST, opts...)
}

func (t *TransactionBuilder) Update(resource interface{}, path string, opts ...BundleEntryOption) *TransactionBuilder {
	return t.addEntry(resource, path, fhir.HTTPVerbPUT, opts...)
}

func (t *TransactionBuilder) addEntry(resource interface{}, path string, verb fhir.HTTPVerb, opts ...BundleEntryOption) *TransactionBuilder {
	data, err := json.Marshal(resource)
	if err != nil {
		return t
	}
	entry := fhir.BundleEntry{
		Resource: data,
		Request: &fhir.BundleEntryRequest{
			Method: verb,
			Url:    path,
		},
	}
	for _, opt := range opts {
		opt(&entry)
	}
	t.Entry = append(t.Entry, entry)
	return t
}

func (t *TransactionBuilder) Bundle() fhir.Bundle {
	return fhir.Bundle(*t)
}

type BundleEntryOption func(entry *fhir.BundleEntry)

func WithFullUrl(fullUrl string) BundleEntryOption {
	return func(entry *fhir.BundleEntry) {
		entry.FullUrl = to.Ptr(fullUrl)
	}
}
