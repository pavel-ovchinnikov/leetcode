delete from Person p1 
where id not in (
    select min(id)
    from Person 
    group by email
)
