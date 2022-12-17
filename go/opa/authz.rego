package authz

# role-permissions assignments

role_permissions := {
	"engineering": [{
		"action": "read",
		"object": "server123",
	}],
	"webdev": [
		{
			"action": "read",
			"object": "server123",
		},
		{
			"action": "write",
			"object": "server123",
		},
	],
	"hr": [{
		"action": "read",
		"object": "database456",
	}],
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
	p == {"action": input.action, "object": input.object}
}
