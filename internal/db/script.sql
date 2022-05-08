CREATE TABLE BookAuthor (
  id int NOT NULL AUTO_INCREMENT,
  author varchar(255),
  book varchar(255),
  PRIMARY KEY (id)
);
INSERT INTO BookAuthor (author, book) VALUES ('King','Green Mile');
INSERT INTO BookAuthor (author, book) VALUES ('Pushkin','Eugeniy Onegin');
INSERT INTO BookAuthor (author, book) VALUES ('Hawking','A Brief History of Time');
