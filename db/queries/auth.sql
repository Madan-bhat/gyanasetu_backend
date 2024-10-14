-- name: UserExists :one
select exists(
    select 1
    from users
    where email = $1
  );
-- name: UserExistsByID :one
select exists(
    select 1
    from users
    where id = $1
  );
-- name: GetGIdByEmail :one 
select gid
from users
where email = $1;
-- name: GetIdByEmail :one
select id
from users
where email = $1;
-- name: CreateUser :one
insert into users (email, name, gid)
values($1, $2, $3)
returning name;
-- name: UpdateRole :exec
update users
set role = $1
where id = $2;
-- name: CreateBDFL :one
insert into users(name, email, gid, bdfl)
values ($1, $2, $3, true)
returning id;
-- name: BDFLExists :one
select exists (
    select 1
    from users
    where bdfl = true
  );
-- name: GetBDFLId :one
select id
from users
where bdfl = true;
-- name: UpdateBDFL :exec
update users
set role = 'bdfl'
where bdfl = true;