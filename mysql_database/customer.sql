use mysql_go;

create table customer
(
    id   varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
);

show create table customer;