-- name: CreateTag :exec
insert into tags (tag_name)
VALUES (?);

-- name: GetLastTag :one
select LAST_INSERT_ID();

-- name: CreateNewTagProduct :exec
INSERT INTO product_tags (product_id, tag_id)
VALUES (?, ?);

-- name: GetProductTags :many
select *
from tags
where tag_id in (select tag_id from product_tags where product_id = ?);

-- name: GetProductTagsID :many
select tag_id
from tags
where tag_id in (select tag_id from product_tags where product_id = ?);

-- name: GetAllTags :many
select *
from tags;

-- name: ExistsTags :one
select exists(select 1 from tags where tag_id = ?);

-- name: GetTagsProduct :many
SELECT c.id      AS commodity_id,
       c.name    AS commodity_name,
       c.price   AS commodity_price,
       (SELECT f.url
        FROM file f
                 INNER JOIN commodity_media cm ON f.id = cm.file_id
        WHERE cm.commodity_id = c.id
        LIMIT 1) AS media_url,
       u.name    AS username,
       u.avatar,
       c.is_free
FROM commodity c
         INNER JOIN
     product_tags pt ON c.id = pt.product_id
         INNER JOIN
     user u ON c.user_id = u.id
WHERE pt.tag_id = ?;

-- name: GetTagName :one
select tag_name
from tags
where tag_id = ?;