#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import numpy as np

print('Derivatives of Erlang B formula')
print('s=? ', end='')
s = int(input())          # positive integer
print('a=? ', end='')
a = float(input())        # offered traffic in erlang

# Calculation of Psi
if a > 1.0:
    # Calculation for pai1
    a1 = 8.5733287401
    a2 = 18.059016973
    a3 = 8.6347608925
    a4 = 0.2677737343
    b1 = 9.5733223454
    b2 = 25.6329561486
    b3 = 21.0996530827
    b4 = 3.9584969228
    
    psn = 1 + a1/a + a2/np.power(a,2) + a3/np.power(a,3) + a4/np.power(a,4)
    psd = 1 + b1/a + b2/np.power(a,2) + b3/np.power(a,3) + b4/np.power(a,4)
    ps1 = psn/psd/a
    
else:
    # Calculation for psi1
    c0 = -0.57721566
    c1 = 0.99999193
    c2 = -0.24991055
    c3 = 0.05519968
    c4 = -9.76004e-3
    c5 = 1.07857e-3
    
    ps1 = np.exp(a) * (c0 + c1*a + c2*np.power(a,2) + c3*np.power(a,3)+c4*np.power(a,4) + c5*np.power(a,5) - np.log10(a))
    
psi = ps1
es = 1.0
for i in range(1, s+1):
    es = a * es/(i + a*es)
    psi = (1 - es)*(psi + 1/i)

# Calculation for derivatives
dea = (s/a - 1 + es)*es
des = -psi*es
dsa = (s/a - 1 + es)/psi

print('s= {0}      a= {1}      Es(a)= {2}'.format(s, a, es))
print('psi = {0}'.format(psi))
print('dEs/da= {0}          dEs/ds= {1}'.format(dea, des))
print('ds/da= {0}'.format(dsa))
    