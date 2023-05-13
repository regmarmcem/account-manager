insert into categories 
    (id, name, deal_type)
values
    (1, '給与収入', 'income'),
    (2, '食費', 'expenditure')
;

insert into deals
    (id, date, amount, category_id, description)
values
    (1, DATE '2023-04-25', 300000, 1, '給料'),
    (2, DATE '2023-04-25', 3000, 2, 'サンエー')
;