# -*-coding:utf-8 -*-
# Author: Eileen Kammel
# Date: 2022-12-18 22:11:24


# -*-coding:utf-8 -*-
# Author: Eileen Kammel
# Date: 2022-12-18 21:41:15

with open('Day1/input.txt', 'r') as f:
    total_inventory = []
    single_inventory = []
    data = f.read().splitlines()
    for food in data:
        if food == '':
            total_inventory.append(single_inventory)
            single_inventory = []
            continue
        single_inventory.append(int(food))
total_inventory = [sum(inventory) for inventory in total_inventory]
total_inventory.sort(reverse=True)
print(sum(total_inventory[:3]))
