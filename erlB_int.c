#include <stdio.h>
#include <math.h>

float Es( int s, float a );

main()
{
  int    s;
  float a, es;

  printf("Erlang B Formula (positive integer server)\n");

  printf("s="); scanf("%d", &s);
  printf("a="); scanf("%g", &a);

  es = Es(s, a);

  printf("s=%d\ta=%g\tEs(a)=%g\n", s, a, es);

}

float Es( int s, float a )
{
  int    i;
  float es=1.0;
  for(i=1; i<=s; i++)
    es = a * es / (i+a*es);

  return(es);

}
