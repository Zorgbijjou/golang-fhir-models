#!/usr/bin/env bash
set -eu

go install

# Fix version to R4 for now.
# As R4B has been released in May 2022, we should update to this or at least
# allow the option to select the version.
wget -O definitions.zip https://www.hl7.org/fhir/R4/definitions.json.zip
unzip definitions.zip profiles-types.json valuesets.json -d fhir
rm definitions.zip
wget -O fhir/bundle.json http://hl7.org/fhir/R4/bundle.profile.json
wget -O fhir/codesystem.json http://hl7.org/fhir/R4/codesystem.profile.json
wget -O fhir/structuredefinition.json http://hl7.org/fhir/R4/structuredefinition.profile.json
wget -O fhir/valueset.json http://hl7.org/fhir/R4/valueset.profile.json

go generate ./fhir
