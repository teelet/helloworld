gcc -shared sotest.c -o libsotest.so

gcc -L. -lsotest main.c -o main

./main