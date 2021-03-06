CREATE TABLE todos (
  id int(11) unsigned not null auto_increment,
  uid varchar(255) not null,
  todo_id int(11) unsigned not null,
  todo varchar(255) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);