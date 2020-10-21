REATE TABLE IF NOT EXISTS users (
    id int(11) unsigned not null auto_increment,
    uid varchar(255) not null,
    email varchar(255),
    name varchar(255),
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    primary key (id)
);

INSERT INTO user (id, uid, email, name) VALUES (1, "0123456789", "test1@test.com", "test user 1");