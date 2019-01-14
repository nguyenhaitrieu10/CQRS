drop table if exists meows;

create table meows {
  id varchar(32) primary key,
  body text not null,
  created_at timestamp with time zone not null
}

