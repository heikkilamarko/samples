package authz

import rego.v1

role_permissions := {
	"sample.admin": ["sample.read", "sample.write"],
	"sample.reader": ["sample.read"],
	"sample.writer": ["sample.read", "sample.write"],
}

default allow := false

allow if {
	p := permissions[_]
	p == input.permission
}

permissions contains p if {
	some r in input.token.resource_access["sample-api"].roles
	some p in role_permissions[r]
}
