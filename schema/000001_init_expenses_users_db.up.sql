CREATE TABLE users
(
    id serial primary key,
    name varchar(256) not null,
    username varchar(256) not null unique,
    password_hash text not null
);

CREATE TABLE expenses
(
    id serial primary key,
    user_id int not null,
    total decimal(10, 2) not null,
    description text,
    category varchar(50) not null,
    created_at timestamp not null default now(),
    foreign key (user_id) references users (id) on delete cascade
);