package routing

type group struct {
	User    user
	Email   email
	File    file
	Product product
	Tags    tags
}

var Group = new(group)
