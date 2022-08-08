-- name: GetTransfer :one
select * from transfers
where id = $1 limit 1;