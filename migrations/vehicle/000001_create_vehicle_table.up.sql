CREATE TABLE IF NOT EXISTS vehicles (
  vin VARCHAR(17) PRIMARY KEY,
  year INT NOT NULL,
  odometer INT NOT NULL,
  msrp NUMERIC,
  price NUMERIC,
  grade NUMERIC,
  brand VARCHAR(50),
  engine VARCHAR(50),
  transmission VARCHAR(50),
  exterior_color VARCHAR(50),
  interior_color VARCHAR(50),
  small_scratches BOOLEAN,
  strong_scratches BOOLEAN,
  electric_fail BOOLEAN,
  suspension_fail BOOLEAN
);