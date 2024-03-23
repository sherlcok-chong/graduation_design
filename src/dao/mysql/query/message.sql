-- name: CreateNewMessage :exec
insert into message (fid, tid, is_file, is_read, texts, createAt)
VALUES (?, ?, ?, ?, ?, ?);

-- name: ReadMessage :exec
update message
set is_read = true
where id = ?;

-- name: GetMessageByUserID :many
select *
from message
where fid = ?;

-- name: GetNotReadMsgByUserID :many
select *
from message
where tid = ?
  and fid = ?
  and is_read = false;

-- name: GetUserWhoTalk :many
select distinct fid
from message
where tid = ?;