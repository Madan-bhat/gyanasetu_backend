create table organizations(
  id serial primary key,
  name varchar(255) not null,
  description text,
  phno varchar(15) [] not null,
  email varchar(255) [] not null,
  address text not null
);