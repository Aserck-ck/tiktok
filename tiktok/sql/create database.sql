use douyin;
drop table users;

create table users(
	id int primary key auto_increment,
	name varchar(255),
    password varchar(255)
);

insert into users values(null,"111","111");
select* from users;

drop table follows;
create table follows(
	id int primary key auto_increment,
    user_id int,
    cancel int default(0),
    foreign key(user_id) references users(id)
);

insert into follows values(null,1,0);

drop table videos;

create table videos(
	id int primary key auto_increment,
    author_id int,
    play_url varchar(255),
    cover_url varchar(255),
    publish_time datetime default(now()),
    title varchar(255),
    foreign key(author_id) references users(id)
);

insert into videos values(null,1,"xx","xx",now(),"xx");
select * from videos;

drop table likes;
create table likes(
	id int primary key auto_increment,
    user_id int,
    video_id int,
    cancel int default(0),
    foreign key(user_id) references users(id),
    foreign key(video_id) references videos(id)
);
select *from likes;
insert into likes values(null,1,1,0);

drop table comments;

create table comments(
	id int primary key auto_increment,
    user_id int,
    video_id int,
    comment_text varchar(255),
    create_date datetime default(now()),
    cancel int default(0),
    foreign key(user_id) references users(id),
    foreign key(video_id) references videos(id)
);
select *from comments;

insert into comments values(null,1,1,"comment",now(),0);