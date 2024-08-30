-- Up Migration
alter table rooms
add column owner_token uuid not null default gen_random_uuid();

-- Down Migration
-- alter table rooms
-- drop column owner_token;