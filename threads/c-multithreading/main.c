#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

#define NUM_THREADS 10

// clang -o exec main.c
// gcc -o exec main.c
// ./exec
void *PrintHello(void *threadid) {
   long tid;
   tid = (long)threadid;
   sleep(1);
   printf("thread %ld: printing\n", tid);
   sleep(15);
   pthread_exit(NULL);
}

int main () {
   pthread_t threads[NUM_THREADS];
   int rc;
   int i;
   for(i = 0; i < NUM_THREADS; i++) {
      printf("main() : creating thread: %d\n", i+1);
      rc = pthread_create(&threads[i], NULL, PrintHello, (void *) (size_t) i+1);
      if (rc) {
         printf("Error: unable to create thread: %d\n", rc);
         exit(-1);
      }
   }

   pthread_exit(NULL);
}
