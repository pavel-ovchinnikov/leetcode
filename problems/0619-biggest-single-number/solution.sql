-- Write your PostgreSQL query statement below

-- ver 1
select max(num) as num
from MyNumbers 
where num in (
    select num
    from MyNumber
    group by num
    having count(num) = 1
)

-- ver 2
select max(num) as num 
from (
    select num
    from MyNumbers
    group by num
    HAVING COUNT(num)=1 
)