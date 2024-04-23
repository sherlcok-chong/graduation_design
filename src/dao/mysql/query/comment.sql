-- name: CreateNewComment :exec
insert into comment (user_id, product_id, texts)
VALUES (?, ?, ?);

-- name: CreateCommentMedias :exec
insert into comment_media (comment_id, file_id)
values (?, ?);
-- name: GetLastCommentID :one
SELECT LAST_INSERT_ID();

-- name: DeleteCommentID :exec
delete
from comment
where id = ?;

-- name: GetProductComment :many
select *
from comment
where product_id = ?;

-- name: GetCommentMedia :many
select file_id
from comment_media
where comment_id = ?;

-- name: GetCommentUser :one
select user_id
from comment
where id = ?;

-- name: DeleteCommentMedia :exec
delete
from comment_media
where comment_id = ?;