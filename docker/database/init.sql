CREATE DATABASE ocp_tenant_api;
CREATE USER ocp_tenant_api_user WITH PASSWORD 'ocp_tenant_api_password';
GRANT all privileges ON DATABASE ocp_tenant_api TO ocp_tenant_api_user;

CREATE TABLE tenant
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    type INT          NOT NULL
);

create index k_tenant_type_name on tenant2 (type, name);

ALTER TABLE tenant OWNER TO ocp_tenant_api;
