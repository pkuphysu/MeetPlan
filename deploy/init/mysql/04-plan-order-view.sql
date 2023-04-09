create or replace view plan_view
as
select plans.*, plans.start_time > now() as is_valid, (plans.quota - count(orders.id))
from plans
         left join orders on plans.id = orders.plan_id and orders.status != 3
group by plans.id;