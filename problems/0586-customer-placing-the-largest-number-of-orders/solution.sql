-- var 1
select customer_number
from Orders 
group by customer_number
order by COUNT(*) DESC
limit 1;


-- var 2
Select customer_number
from Orders
group by customer_number
order by  COUNT(distinct order_number) desc 
limit 1