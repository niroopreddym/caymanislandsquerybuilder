GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;

-- lookup tables and the static data insertion part

create table department(
id int primary key GENERATED ALWAYS AS IDENTITY,
deptname varchar(50)
);

create table role(
id int primary key GENERATED ALWAYS AS IDENTITY,
roletype varchar(50)
);

create table employee (
id int primary key GENERATED ALWAYS AS IDENTITY,
name varchar(255),
deptId int REFERENCES department(id),
role int REFERENCES role(id),
salary int
);
--------------------------------------------------------------
insert into department(deptname) values
('civil'),
('Softwares'),
('Electronics');

insert into role(roletype) values
('developer'),
('sales'),
('tester');

insert into employee(name,deptId,role,salary) values
('A', 2 , 1, 100),
('B', 2 , 1, 200),
('c', 1 , 2, 300),
('D', 3 , 3, 400),
('E', 2 , 3, 150);
