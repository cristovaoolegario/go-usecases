CREATE TABLE orders
(
     id          VARCHAR(255) NOT NULL,
     price       FLOAT NOT NULL,
     tax         FLOAT NOT NULL,
     final_price FLOAT NOT NULL,
     PRIMARY KEY (id)
);