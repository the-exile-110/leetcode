SELECT name
FROM Customers
WHERE id NOT IN (SELECT CustomerId FROM Orders);