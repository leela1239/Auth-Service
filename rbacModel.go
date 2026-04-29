package main

// user --> []role --> []permission

// role1 --> [write , read ,delete ]

// role2 --> [read , write]

type Role string

type Permission string

type RolePermMap map[Role]map[Permission]bool

var RolePerm = RolePermMap{
	"Admin":    map[Permission]bool{"read": true, "delete": true},
	"Operator": map[Permission]bool{"read": true, "write": true},
	"Student":  map[Permission]bool{"read": true},
}

type UserRoleMap map[string][]Role

var UserRole = UserRoleMap{
	"User1": []Role{"Admin"},
	"User2": []Role{"Admin", "Operator"},
	"User3": []Role{"Student"},
}

func IsAuthorized(user string, reqPerm Permission) bool {
	//fetch the roles of the user

	roleList, exists := UserRole[user]
	//check the permissions assigned to each role

	if !exists {
		return false
	}

	for _, role := range roleList {
		if RolePerm[role][reqPerm] {
			return true
		}
	}

	return false
}
