CREATE TABLE tests (
  id int(10) unsigned not null auto_increment PRIMARY KEY,
  test_msg varchar(255) not null
);

INSERT INTO tests (test_msg) values ('hello db')
