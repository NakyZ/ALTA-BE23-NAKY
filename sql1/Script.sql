create table "genre" (
	id int,
	nama varchar (100),
	primary key (id)
);

drop table "user";

create table "students"(
	username varchar (5) primary key,
	"name" varchar (200),
	kartu_pelajar varchar (5),
	hp varchar (13) not null
);

create table "kartu_pelajar"(
	username varchar (5),
	link_foto varchar (200),
	foreign key (username) references students(username),
	primary key (username)
);

create table "genre_buku"(
	id_genre int,
	id_buku int,
	created_at timestamp default now (),
	foreign key (id_genre) references genre (id),
	foreign key (id_buku) references book (id),
	primary key (id_genre, id_buku)
);

insert into genre (nama) values ('thriller');
insert into genre (nama) values ('adventure');
insert into genre (nama) values ('mystery');

insert into book (id,title,penerbit) values 
(1, 'naruto','gramedia'),
(2, 'harry potter','erlangga');

insert into genre-buku (id_genre, id_buku) values 
(1,1)
(5,1)
(3,1)
(4,1)
(5,2)
(5,3);