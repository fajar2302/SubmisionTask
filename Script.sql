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


--4 a
--genres
insert into genres (id, genre_name) values(1,'comedy');
insert into genres (id, genre_name) values(2,'fantasy');
insert into genres (id, genre_name) values(3,'horor');
insert into genres (id, genre_name) values(4,'rommance');
insert into genres(id, genre_name) values(5,'thriller');


--4b
--comedy
insert into books (id, genre_id, title, author, publisher, publish_year) values (1, 1, 'Lintang and the Forbidden Island', 'Tamara Geraldine', 'Bentang Pustaka','2013');
insert into books (id, genre_id, title, author, publisher, publish_year) values (5, 1,'Marmut Merah Jambu','Raditya Dika','Bukune','2019');
insert into books (id, genre_id, title, author, publisher, publish_year) values (6, 1, 'Rindu', 'Tere Liye', 'Republika Penerbit','2020');
insert into books (id, genre_id, title, author, publisher, publish_year) values (7, 1, 'Koala Kumal', 'Raditya Dika', 'Bukune','2015');
insert into books (id, genre_id, title, author, publisher, publish_year) values (8, 1, 'Single', 'Raditya Dika', 'Bukune','2012');


-- fantasy
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (9, 2, 'Harry Potter and the Sorcerer''s Stone', 'J.K. Rowling', 'Scholastic', 1997);
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (10, 2, 'The Hobbit', 'J.R.R. Tolkien', 'George Allen & Unwin', 1937);
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (11, 2, 'Percy Jackson & the Olympians: The Lightning Thief', 'Rick Riordan', 'Miramax Books', 2005);
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (12, 2, 'A Game of Thrones', 'George R.R. Martin', 'Bantam Books', 1996);
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (13, 2, 'The Wise Man''s Fear', 'Patrick Rothfuss', 'DAW Books', 2011);

-- horor
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (14, 3, 'Hantu Jeruk Purut', 'Herman Pratikto', 'GagasMedia', 2009);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (15, 3, 'Tumbal: The Ritual', 'Norman Rantau', 'Gramedia Pustaka Utama', 2019);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (16, 3, 'Pet Sematary', 'Stephen King', 'Doubleday', 1983);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (17, 3, 'Rumah Kertas', 'Tere Liye', 'Gramedia Pustaka Utama', 2006);


--  Romance
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (18, 4, '5 cm', 'Donny Dhirgantoro', 'GagasMedia', 2005);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (19, 4, 'Critical Eleven', 'Ika Natassa', 'Gramedia Pustaka Utama', 2015);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (20, 4, 'Pride and Prejudice', 'Jane Austen', 'Thomas Egerton', 1813);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (21, 4, 'Dear Nathan', 'Erisca Febriani', 'GagasMedia', 2015);


-- Thriller
INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (22, 5, 'The Da Vinci Code', 'Dan Brown', 'Doubleday', 2003);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (23, 5, 'Cinta dalam Kardus', 'Fajar Bustomi', 'GagasMedia', 2011);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (24, 5, 'Gone Girl', 'Gillian Flynn', 'Crown Publishing Group', 2012);

INSERT INTO books (id, genre_id, title, author, publisher, publish_year) 
VALUES (25, 5, 'Hujan Bulan Juni', 'Sapardi Djoko Damono', 'Gramedia Pustaka Utama', 2014);


-- 4c users

-- User 1
INSERT INTO users (id, name, email, password, address, phone, birth_date) 
VALUES (3, 'Budi Santoso', 'budi@example.com', 'password123', 'Jl. Merdeka No. 123, Jakarta', '0812345678901', '1990-05-15');

-- User 2
INSERT INTO users (id, name, email, password, address, phone, birth_date) 
VALUES (4, 'Siti Rahayu', 'siti@example.com', 'password456', 'Jl. Pahlawan No. 456, Surabaya', '0812345678902', '1985-08-20');

-- User 3
INSERT INTO users (id, name, email, password, address, phone, birth_date) 
VALUES (5, 'Agus Widodo', 'agus@example.com', 'password789', 'Jl. Diponegoro No. 789, Bandung', '0812345678903', '1992-12-10');

--4d

SELECT * 
FROM books 
WHERE title = 'Koala Kumal';

--4 e
SELECT * 
FROM books 
WHERE genre_id = (SELECT id FROM genres WHERE genre_name = 'comedy');

--4f

-- Peminjaman 1: User 1 meminjam buku dengan id 1
INSERT INTO loans (id, user_id, book_id, deadline_date, data_loan, status_loan) 
VALUES (1, 1, 1, '2024-06-13', '2024-06-06', 'finished');

-- Peminjaman 2: User 1 meminjam buku dengan id 2
INSERT INTO loans (id, user_id, book_id, deadline_date, data_loan, status_loan) 
VALUES (3, 1, 2, '2024-06-15', CURRENT_DATE, 'unfinished');

-- Peminjaman 3: User 1 meminjam salah satu buku yang bergenre id 2
INSERT INTO loans (id, user_id, book_id, deadline_date, data_loan, status_loan) 
VALUES (4, 1, (SELECT id FROM books WHERE genre_id = 2 LIMIT 1), '2024-06-15', CURRENT_DATE, 'unfinished');

-- Peminjaman 4: User 2 meminjam buku dengan id 3
INSERT INTO loans (id, user_id, book_id, deadline_date, data_loan, status_loan) 
VALUES (5, 2, 3, '2024-06-15', CURRENT_DATE, 'finished');

-- Peminjaman 5: User 3 meminjam buku dengan id 1
INSERT INTO loans (id, user_id, book_id, deadline_date, data_loan, status_loan) 
VALUES (6, 3, 1, '2024-06-15', CURRENT_DATE, 'unfinished');


--4 g
UPDATE loans 
SET status_loan = 'finished'
WHERE id = 3;

--4 h
SELECT u.name as user_name, b.title as book_title
FROM users u
INNER JOIN loans l ON u.id = l.user_id
INNER JOIN books b ON l.book_id = b.id
WHERE l.status_loan = 'unfinished';

--4 i
SELECT u.name as user_name, b.title as book_title
FROM users u
INNER JOIN loans l ON u.id = l.user_id
INNER JOIN books b ON l.book_id = b.id
WHERE l.status_loan = 'finished';


