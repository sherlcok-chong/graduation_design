package logic

type group struct {
	User    user
	Email   email
	File    file
	Product product
	Tags    tags
	Comment comment
}

var Group = new(group)
