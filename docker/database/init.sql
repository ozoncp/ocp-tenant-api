CREATE DATABASE ocp_tenant_api;
CREATE USER ocp_tenant_api_user WITH PASSWORD 'ocp_tenant_api_password';
GRANT all privileges ON DATABASE ocp_tenant_api TO ocp_tenant_api_user;

CREATE TABLE tenant
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    type INT          NOT NULL
);

CREATE INDEX k_tenant_type_name ON tenant (type, name);