create table approvals (
  id serial primary key,
  user_id int references users(id) on delete cascade not null,
  organization_id int references organizations(id) on delete cascade not null,
  status VARCHAR(20) not null default 'pending' check(status in ('pending', 'approved', 'rejected')),
  requested_at timestamp default current_timestamp
);