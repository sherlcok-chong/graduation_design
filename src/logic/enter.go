package logic

type group struct {
	User    user
	Email   email
	File    file
	Product product
	Tags    tags
	Comment comment
	Auto    auto
	Order   order
	WS      ws
}

var Group = new(group)
