SELECT 'CREATE DATABASE spekaspace'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'spekaspace');
CREATE USER postgres WITH ENCRYPTED PASSWORD 'postgres';
GRANT ALL PRIVILEGES ON DATABASE spekaspace TO postgres;
