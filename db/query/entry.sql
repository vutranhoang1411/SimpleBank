-- name: GetEntry :one
select * from entries
where id = $1 limit 1;