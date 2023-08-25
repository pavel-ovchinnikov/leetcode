SELECT
    Users.name,
    SUM(Transactions.amount) AS balance
FROM Users, Transactions
WHERE
    Users.account = Transactions.account
GROUP BY Users.account
HAVING
    balance > 10000;