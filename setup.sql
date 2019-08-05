drop  database if  exists snippetbox;
drop  database if  exists test_snippetbox;
drop  user if exists 'web'@'localhost';
drop  user if exists 'test_web'@'localhost';

create database snippetbox character set utf8mb4 collate utf8mb4_unicode_ci;

create user 'web'@'localhost';
grant select, insert, update on snippetbox.* to 'web'@'localhost';
alter user 'web'@'localhost' identified by 'pass';


create database test_snippetbox character set utf8mb4 collate utf8mb4_unicode_ci;

create user 'test_web'@'localhost';
grant create, drop, alter, index, select, insert, update, delete on test_snippetbox.* to 'test_web'@'localhost';
alter user 'test_web'@'localhost' identified by 'pass';


use snippetbox;

create table snippets (
	id integer not null primary key auto_increment,
	title varchar(100) not null,
	content text not null,
	created datetime not null,
	expires datetime not null
);

create index idx_snippets_created on snippets(created);


insert into snippets (title, content, created, expires) values (
	'An old silent pond',
	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
	utc_timestamp(),
	date_add(utc_timestamp(), interval 365 day)
); 

insert into snippets (title, content, created, expires) values (
	'Over the wintry forest',
	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
	utc_timestamp(),
	date_add(utc_timestamp(), interval 365 day)
);

insert into snippets (title, content, created, expires) values (
	'First autumn morning',
	'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
	utc_timestamp(),
	date_add(utc_timestamp(), interval 7 day)
);

create table users (
	id integer not null primary key auto_increment,
	name varchar(255) not null,
	email varchar(255) not null,
	hashed_password char(60) not null,
	created datetime not null,
	active boolean not null default true
);

alter table users add constraint users_uc_email unique (email);
