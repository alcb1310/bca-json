create table if not exists company (
    id uuid primary key default gen_random_uuid(),
    ruc text not null,
    name text not null,
    employees int not null default 1,

    created_at timestamp with time zone default now(),

    constraint company_name_unique unique (name),
    constraint company_ruc_unique unique (ruc)
);

create table if not exists role (
    id char(2) primary key,
    name text not null,

    created_at timestamp with time zone default now(),

    constraint role_name_unique unique (name)
);

create table if not exists public.user (
    id uuid primary key default gen_random_uuid(),
    email text not null,
    password text not null,
    name text not null,

    created_at timestamp with time zone default now(),
    company_id uuid not null references company (id) on delete restrict,
    role_id char(2) not null references role (id) on delete restrict,

    constraint user_email_unique unique (email)

);
