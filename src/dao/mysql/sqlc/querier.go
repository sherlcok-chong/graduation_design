// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	AddAddressByID(ctx context.Context, arg AddAddressByIDParams) error
	ChangeStatusByOrderID(ctx context.Context, orderID string) error
	CheckUserLike(ctx context.Context, arg CheckUserLikeParams) (bool, error)
	CreateCommentMedias(ctx context.Context, arg CreateCommentMediasParams) error
	CreateFile(ctx context.Context, arg CreateFileParams) error
	CreateNewComment(ctx context.Context, arg CreateNewCommentParams) error
	CreateNewMediaProduct(ctx context.Context, arg CreateNewMediaProductParams) error
	CreateNewMessage(ctx context.Context, arg CreateNewMessageParams) error
	CreateNewTagProduct(ctx context.Context, arg CreateNewTagProductParams) error
	CreateOrder(ctx context.Context, arg CreateOrderParams) error
	CreateProduct(ctx context.Context, arg CreateProductParams) error
	CreateTag(ctx context.Context, tagName string) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteCommentID(ctx context.Context, id int64) error
	DeleteFileByID(ctx context.Context, id int64) error
	DeleteFileMedia(ctx context.Context, commodityID int64) error
	DeleteLike(ctx context.Context, productID int64) error
	DeleteOrder(ctx context.Context, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	DisLikeProduct(ctx context.Context, arg DisLikeProductParams) error
	EnsureExpress(ctx context.Context, id int64) error
	EnsureRec(ctx context.Context, id int64) error
	ExistEmail(ctx context.Context, email string) (bool, error)
	ExistsTags(ctx context.Context, tagID int64) (bool, error)
	ExistsUserByID(ctx context.Context, id int64) (bool, error)
	GetAllTags(ctx context.Context) ([]Tag, error)
	GetCommentMedia(ctx context.Context, commentID int64) ([]int64, error)
	GetCommentUser(ctx context.Context, id int64) (int64, error)
	GetExpiredFileID(ctx context.Context) ([]GetExpiredFileIDRow, error)
	GetFileByID(ctx context.Context, id int64) (string, error)
	GetKeyByID(ctx context.Context, id int64) (string, error)
	GetLastCommentID(ctx context.Context) (int64, error)
	GetLastFileID(ctx context.Context) (int64, error)
	GetLastProductID(ctx context.Context) (int64, error)
	GetLastTag(ctx context.Context) (int64, error)
	GetLikeList(ctx context.Context, userID int64) ([]int64, error)
	GetMessageByUserID(ctx context.Context, fid int64) ([]Message, error)
	GetNotReadMsgByUserID(ctx context.Context, arg GetNotReadMsgByUserIDParams) ([]Message, error)
	GetOrderDetail(ctx context.Context, id int64) (Order, error)
	GetProductByID(ctx context.Context, id int64) (Commodity, error)
	GetProductComment(ctx context.Context, productID int64) ([]Comment, error)
	GetProductFirstMedia(ctx context.Context, commodityID int64) (interface{}, error)
	GetProductInfo(ctx context.Context, offset int32) ([]GetProductInfoRow, error)
	GetProductLike(ctx context.Context, id int64) (GetProductLikeRow, error)
	GetProductMedia(ctx context.Context, id int64) (string, error)
	GetProductMediaId(ctx context.Context, commodityID int64) ([]int64, error)
	GetProductNotFreeTime(ctx context.Context, productID int64) ([]GetProductNotFreeTimeRow, error)
	GetProductTags(ctx context.Context, productID int64) ([]Tag, error)
	GetTagName(ctx context.Context, tagID int64) (string, error)
	GetTagsProduct(ctx context.Context, tagID int64) ([]GetTagsProductRow, error)
	GetUserAvatar(ctx context.Context, arg GetUserAvatarParams) (string, error)
	GetUserAvatarByID(ctx context.Context, id int64) (string, error)
	GetUserBorrowOrder(ctx context.Context, lendUserID int64) ([]Order, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, name string) (User, error)
	GetUserInfoById(ctx context.Context, id int64) (User, error)
	GetUserLendOrder(ctx context.Context, borrowUserID int64) ([]Order, error)
	GetUserLendProduct(ctx context.Context, userID int64) ([]Commodity, error)
	GetUserNeedInfo(ctx context.Context, userID int64) ([]GetUserNeedInfoRow, error)
	GetUserProductInfo(ctx context.Context, userID int64) ([]GetUserProductInfoRow, error)
	GetUserWhoTalk(ctx context.Context, tid int64) ([]int64, error)
	LikeProduct(ctx context.Context, arg LikeProductParams) error
	ReadMessage(ctx context.Context, id int64) error
	ReadUserMessage(ctx context.Context, arg ReadUserMessageParams) error
	SearchLikeText(ctx context.Context, concat interface{}) ([]SearchLikeTextRow, error)
	UpdateOrderExpress(ctx context.Context, arg UpdateOrderExpressParams) error
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	UpdateUserAvatar(ctx context.Context, arg UpdateUserAvatarParams) error
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error
}

var _ Querier = (*Queries)(nil)
