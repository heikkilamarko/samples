ALTER SESSION SET CONTAINER=FREEPDB1;

CREATE USER DEMO IDENTIFIED BY demopwd QUOTA UNLIMITED ON USERS;

GRANT CONNECT, RESOURCE TO DEMO;
