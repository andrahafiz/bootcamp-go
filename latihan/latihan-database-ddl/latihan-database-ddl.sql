create database toko_online_noobee;

create table users(
	user_id serial primary key,
	username varchar(50) unique not null,
	email varchar(50) unique not null,
	nama_lengkap varchar(100) not null
);

create table orders(
	order_id serial primary key,
	user_id int,
	tanggal_pemesanan timestamp default now(),
	status varchar(50) not null
);

alter table orders 
add foreign key (user_id) references users(user_id);


create table order_items(
 	item_id serial primary key,
	order_id int,
	product_name varchar(50) not null,
	quantity int not null,
	harga_per_item int not null
);

alter table order_items 
add foreign key (order_id) references orders(order_id);