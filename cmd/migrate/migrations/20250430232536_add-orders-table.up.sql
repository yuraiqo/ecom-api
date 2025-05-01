CREATE TYPE order_status AS ENUM('pending', 'completed', 'canceled');

CREATE TABLE IF NOT EXISTS orders (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "userId" INT NOT NULL,
    "total" DECIMAL(10, 2) NOT NULL,
    "status" order_status NOT NULL DEFAULT 'pending',
    "address" TEXT NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("userId") REFERENCES users("id")
);