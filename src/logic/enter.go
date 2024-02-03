package logic

type group struct {
	User    user
	Email   email
	File    file
	Product product
}

var Group = new(group)
