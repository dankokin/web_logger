CREATE TABLE if not exists Messages (
    target varchar(256),
    request_uri varchar(256),
    status_code int,
    request_rules_check_elapsed bigint,
    response_rules_check_elapsed bigint,
    http_elapsed bigint,
    request_size bigint,
    response_size bigint,
    registered_at TIMESTAMP WITH TIME ZONE
);
