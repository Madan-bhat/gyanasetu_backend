CREATE TABLE users (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email text NOT NULL UNIQUE,
  gid text NOT NULL UNIQUE,
  organization_id int REFERENCES organizations(id) ON DELETE CASCADE,
  role varchar(10) CHECK(role IN ('teacher', 'student', 'bdfl')),
  organization_role varchar(10) CHECK(
    organization_role IN('superadmin', 'admin', 'member')
  ),
  bdfl boolean DEFAULT false unique
);
create unique index unique_bdfl_true on users(bdfl)
where bdfl = true;