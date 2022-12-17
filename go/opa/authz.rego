package authz

import future.keywords.contains
import future.keywords.in

role_permissions := {
	"sample.admin": ["sample.read", "sample.write"],
	"sample.reader": ["sample.read"],
	"sample.writer": ["sample.read", "sample.write"],
}

default allow = false

allow {
	p := permissions[_]
	p == input.permission
}

permissions contains p {
	some r in input.token.resource_access["sample-api"].roles
	some p in role_permissions[r]
}
