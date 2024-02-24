-- name: CreateFile :exec
insert into file (filename, file_key, url, userid)
values (?, ?, ?, ?);

-- name: GetLastFileID :one
SELECT LAST_INSERT_ID();

-- name: GetUserAvatar :one
select url
from file
where userid = ?
  and filename = ?;

-- name: GetFileByID :one
select url
from file
where id = ?;

-- name: GetKeyByID :one
select file_key
from file
where id = ?;

-- name: DeleteFileByID :exec
delete from file where id = ?;