-- Write your PostgreSQL query statement below

-- var 1
select cur_day.id
from Weather as cur_day
where exists (
    select 1
    from Weather as yesterday
    where cur_day.temperature > yesterday.temperature
    and cur_day.recordDate = yesterday.recordDate + 1
)


-- var 2
SELECT w1.id
FROM Weather w1
JOIN Weather w2 ON w1.recordDate = w2.recordDate + 1
WHERE w1.temperature > w2.temperature;