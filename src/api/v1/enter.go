package v1

type group struct {
	User    user
	Email   email
	File    file
	Product product
	Tags    tags
	Comment comment
	Orders  orders
	Alipay  alipay
	Ws      ws
}

var Group = new(group)
