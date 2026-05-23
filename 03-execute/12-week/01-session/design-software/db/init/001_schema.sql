create table if not exists instructors (
  id bigserial primary key,
  name text not null unique,
  instructor_type text not null check (instructor_type in ('planta', 'contratista')),
  created_at timestamptz not null default now()
);

create table if not exists training_programs (
  id bigserial primary key,
  name text not null unique,
  level text,
  created_at timestamptz not null default now()
);

create table if not exists fichas (
  id bigserial primary key,
  code text not null unique,
  program_id bigint references training_programs(id),
  created_at timestamptz not null default now()
);

create table if not exists training_environments (
  id bigserial primary key,
  name text not null unique,
  created_at timestamptz not null default now()
);

create table if not exists schedule_entries (
  id bigserial primary key,
  date date not null,
  start_time time not null,
  end_time time not null,
  instructor_id bigint not null references instructors(id),
  ficha_id bigint not null references fichas(id),
  environment_id bigint not null references training_environments(id),
  program_id bigint references training_programs(id),
  competence text,
  rap text,
  learning_activity text,
  observation text,
  created_at timestamptz not null default now(),
  check (end_time > start_time)
);

create index if not exists idx_schedule_instructor_time on schedule_entries (date, instructor_id, start_time, end_time);
create index if not exists idx_schedule_ficha_time on schedule_entries (date, ficha_id, start_time, end_time);
create index if not exists idx_schedule_environment_time on schedule_entries (date, environment_id, start_time, end_time);

create table if not exists no_programming_days (
  id bigserial primary key,
  date date not null,
  ficha_id bigint references fichas(id),
  instructor_id bigint references instructors(id),
  reason text not null,
  created_at timestamptz not null default now()
);

