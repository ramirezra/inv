CREATE DATABASE leads;

\c leads

CREATE USER leads WITH PASSWORD 'XXXXX';

GRANT ALL PRIVILEGES ON DATABASE leads to leads;

GRANT ALL PRIVILEGES ON TABLE leads to leads;

CREATE TABLE leads (
   ID       INT PRIMARY KEY     NOT NULL,
   STATUS   TEXT NOT NULL,
   CONTACT  TEXT NOT NULL,
   SALES        TEXT    NOT NULL,
   VALUE FLOAT NOT NULL
);

INSERT INTO leads (id, status, contact, sales, value) VALUES
('1','Lead','Devilbiss','Arie','20000.00'),
('2','Negotiation','Phillips','Lynda','45000.00'),
('3','Proposal','Phillips','Lynda','20000.00'),
('4','Lead','MSA','Deb','1500.00');
