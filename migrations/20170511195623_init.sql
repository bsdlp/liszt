-- +goose Up
begin;

create table residents (
	id int not null auto_increment,
	firstname varchar(255),
	middlename varchar(255),
	lastname varchar(255),
	primary key (id)
);

create table units (
	id int not null auto_increment,
	name varchar(100),
	primary key (id)
);

create table units_residents (
	unit int not null references units (id),
	resident int not null references residents (id),
	primary key (unit, resident)
);

commit;

-- +goose Down
begin;

drop table units_residents;
drop table units;
drop table residents;

commit;
