create table todo (
    id int identity constraint todo_pk primary key,
    name nvarchar(100) not null,
    description nvarchar(1000)
) create unique index todo_name_uindex on todo (name);