insert  into users (username, email, nama_lengkap) values
('user1','user1@example.com','User satu'),
('user2','user2@example.com','User dua'),
('user3','user3@example.com','User tiga');

insert  into orders values
('1','1','2023-10-04 20:59:50.815398','Dalam Pengiriman'),
('2','2','2023-10-04 20:59:50.815398','Selesai'),
('3','3','2023-10-04 21:59:50.815398','Dibatalkan');

insert  into order_items values
('1','1','Produk A','2','50000'),
('2','1','Produk B','3','30000'),
('3','2','Produk C','1','75000'),
('4','2','Produk D','2','60000');

select o.order_id, u.username
		from orders as o
		join users as u on o.user_id = u.user_id;
	
select o.order_id, u.username,sum(oi.harga_per_item*oi.quantity)  
		from orders as o
		join users as u on o.user_id = u.user_id
		join order_items as oi on o.order_id = oi.order_id
		GROUP BY o.order_id, u.username;
	

