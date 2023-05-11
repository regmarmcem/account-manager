insert into categories 
    (id, name, deal_type)
values
    (1, '給与収入', 'income'),
    (2, '食費', 'expenditure')
;

insert into deals
    (id, date, category_id, description)
values
    (1, DATE '2023-04-25', 1, '給料'),
    (2, DATE '2023-04-25', 2, 'サンエー')
;