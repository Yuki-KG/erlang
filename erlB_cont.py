#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import numpy as np

print('Erlang B Formula (continuous server)')

print('s=? ', end='')
s = float(input())      # integer number of servers
print('a=? ', end='')
a = float(input())      # offered traffic in erlang

# Calculation for x and k (truncation terms)
n = int(s)
x = s - n
k = int(5/4*np.sqrt(x+500) + 4/a)

# Calculation for continued fraction
esk = a
for i in range(k, 0, -1):
    esk = a + (-x + i - 1)/(1 + i/esk)
ex = esk/a
    
# Calculation for Es(a)
es = ex
for j in range(1, n+1):
    es = a * es/(x + j + a*es)


print('s= {0}     a= {1}     Es(a)= {2}'.format(s, a, es))