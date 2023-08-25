-- Write your MySQL query statement below

UPDATE Salary 
SET sex = (CASE WHEN sex = 'f' THEN 'm' ELSE 'f' END) 

-- update Salary set sex= if(sex='f','m','f');