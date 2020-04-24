# Grocey specs

We want to keep it simple


orders table

| sugar | salt | bread | username | date |
|1      | 12    | 12 | adonese  | 121 |


`
create table products (
    name text
    product_id integer
)

create table prices (
    id int
    unit_price real
)

create table carts (
	id integer primary key,
	user_id integer,
	created_at time,
	delivery time,
	is_completed bool,
	product_id integer,
	quantity integer,
	token text
);

create table cartitems (
    id integer primary key,
    product_it integer,
    cart_id integer,
    user_id integer
)

[Cart]
CartId		int	IDENTITY Primary Key Not Null
UserId		int
DateCreated	datetime NOT NULL
CheckedOut	bit NOT NULL Default 0

[CartItems]
CartItemId	int	IDENTITY Primary Key Not Null
CartId		int
ProductId	int
Quantity	int
Price		money

`