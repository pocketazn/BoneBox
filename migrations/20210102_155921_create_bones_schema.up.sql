create schema if not exists bonebox;

create table if not exists bonebox.bones(
    id bigint not null generated always as identity primary key,
    name text not null, -- short name
    description text, -- lengthy description
    external_label text, -- if an external source has their own identifier they want to associate with the bone for auditing
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    archived_at timestamptz
    );