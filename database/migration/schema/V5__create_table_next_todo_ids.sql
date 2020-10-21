CREATE TABLE next_todo_ids (
  id int(11) unsigned not null auto_increment,
  uid varchar(255) not null,
  next_todo_id int(11) unsigned not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);
