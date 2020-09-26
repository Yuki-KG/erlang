#!/usr/bin/env python3
# -*- coding: utf-8 -*-

print('Erlang B Formula (positive integer server)')

print('s=? ', end='')
s = int(input())        # integer number of servers
print('a=? ', end='')
a = float(input())      # offered traffic in erlang

# Calculation of Es(a)
es = 1.0
for i in range(1, s+1):
    es = a * es/(i+a*es)


print('s= {0}     a= {1}     Es(a)= {2}'.format(s, a, es))