use
    mysql_go;

select *
from customer;

desc customer;

delete
from customer;

alter table customer
    add column email      varchar(100),
    add column balance    int       default 0,
    add column rating     double    default 0.0,
    add column created_at timestamp default current_timestamp,
    add column birth_date date,
    add column married    boolean   default false;


insert into customer(id, name, balance, rating, married)
values ('joko', 'Joko', 50000, 4.5, true);

create table authentification
(
    username varchar(100) not null,
    password varchar(100) not null,
    primary key (username)
);

select *
from authentification;

insert into authentification (username, password)
values ('admin', 'admin');

create table comment
(
    id      int          not null auto_increment,
    email   varchar(100) not null,
    comment text,
    primary key (id)
);

desc comment;

select *
from comment;

select count(*)
from comment;