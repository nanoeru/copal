#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int fibo(int n) {
	if (n < 2) {
		return n;
	} else {
		return fibo(n - 1) + fibo(n - 2);
	}
}

int main(int args, char *argv[]) {
	clock_t tc0, tc1;
	int x;
	int n = 30;
	tc0 = clock();
	x = fibo(n);
	tc1 = clock();
	printf("%f\n",((double)tc1 - tc0)/CLOCKS_PER_SEC);
	printf("%d\n", x);
	return 0;
}
