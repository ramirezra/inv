CREATE TABLE po_headers (
  loc CHAR(2) NOT NULL,
  po_id SERIAL PRIMARY KEY  NOT NULL,
  -- PRIMARY KEY (loc, po_id),
  po_number VARCHAR,
  vendor_no CHAR(6) NOT NULL
  );

ALTER SEQUENCE po_headers_po_id_seq RESTART WITH 600000;

CREATE TABLE po_lines (
  po_line_id SERIAL PRIMARY KEY NOT NULL,
  po_id INT REFERENCES po_headers (po_id),
  prc CHAR(3),
  partno VARCHAR(50),
  qty INT,
  cost FLOAT
);

INSERT INTO po_headers
(loc,vendor_no)
VALUES('10','75000');

-- INSERT INTO po_headers (po_number) SELECT (loc,po_id) FROM po_headers;

INSERT INTO po_lines
(prc,partno,qty,cost)
VALUES ('SLP', 'MENB1060AF03', 2000, 8.98);
