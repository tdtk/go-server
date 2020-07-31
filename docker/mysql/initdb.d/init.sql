CREATE TABLE role
(
  role_id int,
  role_name varchar(50)
);

insert into role
values
  (1, "管理者");
insert into role
values
  (2, "一般");

CREATE TABLE user_info
(
  user_id serial auto_increment not null primary key ,
  login_id varchar(50) unique,
  user_name varchar(50),
  telephone varchar(50),
  password varchar(50),
  role_id int,
  is_deleted boolean
);

insert into user_info
values
  (null, "sato", "佐藤", "0901111111", "axiz", 1, 0);
insert into user_info
values
  (null, "suzuki", "鈴木", "09022222222", "bxiz", 1, 0);
insert into user_info
values
  (null, "takahashi", "高橋", "09033333333", "cxiz", 1, 0);
insert into user_info 
values
  (null, "tanaka", "田中", "09044444444", "dxiz", 2, 0);
insert into user_info 
values
  (null, "ito", "伊藤", "09055555555", "exiz", 2, 0);
insert into user_info 
values
  (null, "yamamoto", "山本", "09066666666", "fxiz", 2, 0);
insert into user_info 
values
  (null, "watanabe", "渡辺", "09077777777", "gxiz", 2, 0);
insert into user_info 
values
  (null, "nakamura", "中村", "09088888888", "hxiz", 2, 0);
insert into user_info 
values
  (null, "kobayashi", "小林", "09099999999", "ixiz", 2, 0);
insert into user_info 
values
  (null, "kato", "加藤", "09000000000", "jxiz", 2, 0);