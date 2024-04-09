package routing

type group struct {
	User    user
	Email   email
	File    file
	Product product
	Tags    tags
	Comment comment
	Msg     msg
	Order   order
}

var Group = new(group)
