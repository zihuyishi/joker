create table public.joker
(
	id bigserial not null
		constraint joker_pk
			primary key,
	title varchar(100) not null,
	content text not null,
	time timestamp default CURRENT_TIMESTAMP not null
);

alter table public.joker owner to postgres;

create table public.tag
(
	id bigserial not null
		constraint tag_pk
			primary key,
	name varchar(50) not null
);

alter table public.tag owner to postgres;

create unique index tag_name_uindex
	on public.tag (name);

create table public.joker_tag
(
	id bigserial not null
		constraint joker_tag_pk
			primary key,
	joker_id bigserial not null
		constraint joker_tag_joker_id_fk
			references public.joker
				on delete cascade,
	tag_id bigserial not null
		constraint joker_tag_tag_id_fk
			references public.tag
				on delete cascade,
	constraint joker_tag_pk_2
		unique (joker_id, tag_id)
);

alter table public.joker_tag owner to postgres;

create table public."user"
(
	id bigserial not null
		constraint user_pk
			primary key,
	name varchar(50) not null,
	password varchar(255) not null,
	time timestamp default CURRENT_TIMESTAMP not null
);

alter table public."user" owner to postgres;

create unique index user_name_uindex
	on public."user" (name);

