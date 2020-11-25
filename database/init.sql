CREATE TABLE if not exists Messages (
    registered_at TIMESTAMP WITH TIME ZONE,
    message varchar(256),
    target varchar(256)
);
