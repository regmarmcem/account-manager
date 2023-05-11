CREATE TYPE deal_type AS ENUM ('income', 'expenditure');

CREATE TABLE categories(
    id int PRIMARY KEY,
    name varchar(100),
    deal_type deal_type
);

CREATE TABLE deals(
    id int,
    date date,
    amount int,
    category_id int,
    description text,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);