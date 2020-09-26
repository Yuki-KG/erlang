#include <stdio.h>
#include <math.h>

float Xk( float s, float a, int *k, int *n );
float Ex( float s, float a, int k, float x, int n );
float Es(float a, float ex, float x, int n);

main()
{
  int k, n;
  float s, a, x, ex, es;

  printf("Erlang B Formula (continuous server)\n");

  printf("s="); scanf("%f", &s);
  printf("a="); scanf("%f", &a);

  x = Xk(s,a, &k, &n);
  ex = Ex(s,a,k,x,n);
  es = Es(a,ex,x,n);

  printf("s=%f\ta=%f\tEs(a)=%f\n", s, a, es);
  
}

float Xk( float s, float a, int *k, int *n )
{
  float x;
  *n = (int)s;  x = s - *n;
  *k = (int)(5/4*sqrt(x+500) + 4/a);
  return(x);
}

float Ex( float s, float a, int k, float x, int n )
{
  float esk;
  int   i;
  float ex;

  esk = a;
  for(i=k; i>=1; i--)
    esk = a + (-x + i - 1)/(1 + i/esk);
  ex = esk/a;
  return(ex);
}

float Es(float a, float ex, float x, int n)
{
  int   j;
  float es;
  es = ex;
  for(j=1; j<=n; j++)
    es = a * es/(x+j+a*es);
  return(es);
}
