drop table loans;
drop table books;
drop table genres;

create table "genres"(
	id serial primary key,
	name varchar(200) not null,
	created_at timestamp default now()
);

create table books(
	id serial primary key,
	title varchar(200) not null,
	author varchar(200),
	publisher varchar(200),
	publish_year int,
	genre_id int,
	created_at timestamp default now()
);

create table loans(
	id serial primary key,
	user_id int,
	book_id int,
	deadline date not null,
	"return" date default now(),
	status bool default true,
	created_at timestamp default now(),
	foreign key (user_id) references users(id),
	foreign key (book_id) references books(id)
);

alter table books
add foreign key(genre_id) references genres(id);

insert into genres(name) values 
('horor'),
('comedy'),
('drama'),
('thriller'),
('action'),
('romance'),
('mystery');

insert into users(name, email, "password", phone, dob) values 
('jerry', 'jerry@alterra.id', 'p12345', '12345', '2021-09-21');
insert into users(name, email, "password", phone, dob) values 
('kamil', 'kamil@alterra.id', 'p12345', '12345', '2021-09-21');
insert into users(name, email, "password", phone, dob) values 
('fakhry', 'fakhry@alterra.id', 'p12345', '12345', '2021-09-21');
insert into users(name, email, "password", phone, dob) values 
('sakinah', 'sakinah@alterra.id', 'p12345', '12345', '2021-09-21');
insert into users(name, email, "password", phone, dob) values 
('ruth', 'ruth@alterra.id', 'p12345', '12345', '2021-09-21');

insert into books(title, author, publisher, publish_year, genre_id)
values ('harry potter 1', 'JK. Rowling', 'Gramedia', 2010, 7);
insert into books(title, author, publish_year, genre_id)
values ('harry potter 2','JK. Rowling', 2010, 7);
insert into books(title, author, publish_year, genre_id)
values ('harry potter 3','JK. Rowling', 2010, 7);

insert into loans(user_id, book_id, deadline) values 
(1, 2, '2024-06-10'),
(2, 3, '2024-06-10');
insert into loans(user_id, book_id, deadline) values 
(1, 4, '2024-06-10');


select * 
from users;

select * 
from users
limit 2;

select name, phone, dob
from users;

select name, phone, dob
from users;

select name "nama", phone as "telepon", dob as "tanggal lahir"
from users;

delete from users where id=3;

select name "nama", phone as "telepon", dob as "tanggal lahir"
from users
where id = 2;
select name "nama", phone as "telepon", dob as "tanggal lahir"
from users
where id > 2;

select name "nama", phone as "telepon", dob as "tanggal lahir"
from users
where name like 'j%';

select name "nama", phone as "telepon", dob as "tanggal lahir"
from users
where name like '%a%' and id > 3
order by phone, name desc;

select name "nama", phone as "telepon", dob as "tanggal lahir"
from users
order by phone desc, name desc;

-- Tampilkan nama buku beserta genrenya

select b.title , g."name" 
from books b
join genres g on g.id = b.id;

-- Tampilkan nama orang beserta judul buku dan genre buku yang dia pinjam

select u."name" , b.title , g."name" 
from users u
join loans l on l.user_id = u.id 
join books b on b.id = l.book_id 
join genres g on g.id =b.genre_id; 

select u."name" , b.title , g."name" 
from loans l
join users u on u.id = l.user_id 
join books b on b.id = l.book_id 
join genres g on g.id =b.genre_id;

select u."name" , b.title , g."name" 
from loans l
left join users u on u.id = l.user_id 
left join books b on b.id = l.book_id 
left join genres g on g.id =b.genre_id;

select * from loans;
select * from users;

select u."name", l.id 
from loans l
right join users u on u.id = l.user_id;

select u."name", l.id 
from users u
left join loans l on u.id = l.user_id;

select u."name" , b.title , g."name" 
from loans l
right join users u on u.id = l.user_id 
right join books b on b.id = l.book_id 
right join genres g on g.id =b.genre_id;

select u."name" , b.title , g."name" 
from loans l
left join users u on u.id = l.user_id 
right join books b on b.id = l.book_id 
right join genres g on g.id =b.genre_id;


-- Tampilkan data semua orang, beserta judul buku yang pernah dia pinjam.

select u."name" , b.title
from loans l
right join users u on u.id = l.user_id 
left join books b on b.id = l.book_id ;