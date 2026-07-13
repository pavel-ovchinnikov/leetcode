select e.name as Employee
from Employee e
inner join Employee m on e.ManagerID is not NULL 
and e.ManagerId = m.Id 
and e.Salary > m.Salary