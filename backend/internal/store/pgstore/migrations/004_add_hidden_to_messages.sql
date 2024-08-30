-- Up Migration
alter table messages
add column hidden boolean not null default false;

-- Down Migration
-- alter table messages
-- drop column hidden;