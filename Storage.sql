CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

CREATE database my_db;
USE my_db;

-- Create of Passenger & Driver Information
CREATE TABLE Passenger (PasFirstName varchar(20), PasLastName VARCHAR(20), PasMobileNum int, PasEmailAdd varchar (20));


CREATE TABLE Driver (FirstName varchar (20) NOT NULL PRIMARY KEY,
LastName VARCHAR(20), MobileNum int, EmailAdd varchar (20), IdNum int, CarLicen int);


-- Enter one Passenger & Driver Information
INSERT INTO Passenger VALUES ('Nigel', 'Lim', 91234567, 'nigel@gmail.com');
INSERT INTO Driver VALUES ('Sam', 'Tan', 81234567, 'sam@gmail.com', 9907856, 1234);

DROP TABLE Passenger;
DROP TABLE Driver;