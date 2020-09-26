#include <stdio.h>
#include <math.h>


double Psi(int s, float a, double *es);
double Ps1a(float a);
double Ps1c(float a);
void Der(int s, float a, double es, double psi, double *dea, double *des,
	   double *dsa);

main()
{
  int s;
  float a;
  double es, psi, psia, dea, des, dsa;

  printf("Derivations of Erlang B formula\n");
  printf("s="); scanf("%d", &s);
  printf("a="); scanf("%f", &a);

  psi = Psi(s,a, &es);
  Der(s,a,es,psi, &dea, &des, &dsa);

  printf("s=%d\ta=%f\tEs(a)=%g\n", s, a, es);
  printf("¦×=%g\n", psi);
  printf("dEs/da=%g\tdEs/ds=%g\n", dea, des);
  printf("ds/da=%g\n", dsa);

}

double Psi(int s, float a, double *es)
{
  int i;
  double psi, ps1;

  if(a>1) ps1=Ps1a(a); else ps1=Ps1c(a);
  psi = ps1; *es = 1;
  for(i=1; i<=s; i++){
    *es = a * (*es)/(i+a*(*es));  psi = (1 - *es) * (psi+1/i);
  }
  return(psi);
}

double Ps1a(float a)
{
  double a1 = 8.5733287401; double a2 = 18.059016973;
  double a3 = 8.6347608925; double a4 = 0.2677737343;
  double b1 = 9.5733223454; double b2 = 25.6329561486;
  double b3 = 21.0996530827; double b4 = 3.9584969228;
  double psn, psd, ps1;

  psn = 1 + a1/a + a2/pow(a,2) + a3/pow(a,3) + a4/pow(a,4);
  psd = 1 + b1/a + b2/pow(a,2) + b3/pow(a,3) + b4/pow(a,4);
  ps1 = psn/psd/a;
  return(ps1);
}

double Ps1c(float a)
{
  double c0 = -0.57721566; double c1 = 0.99999193;
  double c2 = -0.24991055; double c3 = 0.05519968;
  double c4 = -9.76004e-3; double c5 = 1.07857e-3;
  double ps1;

  ps1 = exp(a)*(c0 + c1*a + c2*pow(a,2) + c3*pow(a,3) + c4*pow(a,4)
		+ c5*pow(a,5) - log10(a));
  return(ps1);
}

void Der(int s, float a, double es, double psi, double *dea, double *des,
	   double *dsa)
{
  *dea = ((double)s/a - 1 + es)*es;
  *des = -psi * es;
  *dsa = ((double)s/a - 1 + es)/psi;
}

