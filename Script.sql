create table "users"(
	id int,
	name varchar(100),
	email varchar(100),
	password varchar(50),
	address varchar(100),
	phone varchar(13),
	birth_date date,
	primary key (id)
);

create table "genres"(
	id int,
	genre_name varchar(50),
	primary key (id)
);

create table "books"(
	id int,
	genre_id int,
	title varchar(50),
	author varchar(50),
	publisher varchar(50),
	publish_year varchar(4),
	primary key (id),
	foreign key (genre_id) references genres(id)
);


create table "loans"(
	id int,
	user_id int,
	book_id int,
	deadline_date date,
	data_loan date,
	data_return date,
	status_loan varchar(20),
	primary key (id),
	foreign key (user_id) references users(id),
	foreign key (book_id) references books(id)
);

create table "admins"(
	id int, 
	name varchar(50),
	email varchar(50),
	password varchar(50),
	phone varchar(13),
	primary key (id)
);
create table "book_request"(
	id int,
	user_id int,
	approved_admin_id int,
	title varchar(50),
	author varchar(50),
	publisher varchar(50),
	publish varchar(50),
	status_request varchar(20),
	primary key (id),
	foreign key (user_id) references users(id),
	foreign key (approved_admin_id) references admins(id)
);



