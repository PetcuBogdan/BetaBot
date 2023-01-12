--CREATE TABLE TASK ( id NUMBER PRIMARY KEY, taskName NOT NULL, shop NOT NULL, productName NOT NULL,
--category NOT NULL, size, color, profile NOT NULL, card NOT NULL, status);

--INSERT INTO task(id, taskName, shop, productName, category, size,  color, profile, card, status)
--VALUES(1, "task1", "Nike", "Jordan1", "Shoe", "43","blue", "profile1", "card1", "paused");

--SELECT * FROM task;

--CREATE TABLE PROFILE ( id NUMBER PRIMARY KEY,  fname NOT NULL, lname NOT NULL, email NOT NULL,
--phone NOT NULL, address NOT NULL, address2, postcode NOT NULL, city NOT NULL, county NOT NULL, country NOT NULL);

--INSERT INTO profile(id, fname, lname, email, phone, address, address2, postcode, city, county, country)
--VALUES(1,"test","test","test","test","test","test","test","test","test","test");

--SELECT * FROM profile;

--CREATE TABLE card ( id NUMBER PRIMARY KEY, name NOT NULL, nameCard NOT NULL,
--cardNumber NOT NULL, expirationDate NOT NULL, cvv NOT NULL);

--INSERT INTO card (id, name, nameCard, cardNumber, expirationDate, cvv)
--VALUES(1, "test", "test", "test", "test", "test");

--DELETE FROM profile WHERE id=3

--SELECT MAX(id) FROM card
--SELECT * FROM card

--CREATE TABLE IF NOT EXISTS card (id INTEGER PRIMARY KEY, nameCard TEXT, cardNumber TEXT, expirationDate TEXT, cvv TEXT)

--CREATE  TABLE supreme (productName TEXT NOT NUll, category TEXT NOT NULL, colors TEXT NOT NULL)

--INSERT INTO supreme(productName, category, colors)
--VALUES("GORE-TEX Camp Cap", "hats", "Silver Olive Blue Black")

--delete from supreme where productName='GOTE-TEX Camp Cap'