-- name: CreateTag :exec
insert into tags (tag_name)
VALUES (?);

-- name: GetLastTag :one
select LAST_INSERT_ID();

-- name: CreateNewTagProduct :exec
INSERT INTO product_tags (product_id, tag_id)
VALUES (?, ?);

-- name: GetProductTags :many
select tag_name
from tags
where tag_id in (select tag_id from product_tags where product_id = ?);
