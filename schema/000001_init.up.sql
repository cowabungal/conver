CREATE TABLE users
(
    id       serial       not null unique,
    username varchar(255) not null,
    user_id varchar (30) not null unique,
    state varchar(30) not null default 'default'
);

CREATE TABLE callbacks
(
    id       serial       not null unique,
    user_id varchar (30) not null unique,
    callback_id varchar (30) not null unique,
    callback_data varchar (30)
)