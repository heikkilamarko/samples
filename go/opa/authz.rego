package authz

# role-permissions assignments

role_permissions := {
	"sample.admin": ["sample.read", "sample.write"],
	"sample.reader": ["sample.read"],
	"sample.writer": ["sample.read", "sample.write"],
}

# RBAC logic

default allow = false

allow {
	# lookup the list of roles for the token
	roles := input.token.resource_access["sample-api"].roles

	# for each role in that list
	r := roles[_]

	# lookup the permissions list for role r
	permissions := role_permissions[r]

	# for each permission
	p := permissions[_]

	# check if the permission granted to r matches the user's request
	p == input.permission
}
