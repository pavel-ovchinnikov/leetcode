
--var 1
select s.name
from SalesPerson  s
where s.sales_id NOT IN (
    select s.sales_id
    from Orders o
    inner join Company c on c.com_id = o.com_id
    where o.sales_id = s.sales_id and c.name = 'RED'
)


-- var 2
select name 
from salesperson
where sales_id not in (
    select e.sales_id
    from salesperson as e
    left join orders as o on o.sales_id = e.sales_id
    left join company as c on o.com_id = c.com_id
    where c.name = 'RED'
)