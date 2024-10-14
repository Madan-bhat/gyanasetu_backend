-- name: GetOrgIDByName :one
select id
from organizations
where name = $1;
-- name: GetAllOrganizations :many
select *
from organizations;
-- name: CreateOrganization :one
INSERT into organizations (name, description, phno, email, address)
values ($1, $2, $3, $4, $5)
returning name;
-- name: OrgExistsByName :one
select exists (
    select 1
    from organizations
    where name ilike $1
  );
-- name: CreateApproval :exec
insert into approvals (user_id, organization_id)
values ($1, $2);
-- name: IsBDFL :one
select bdfl
from users
where id = $1;