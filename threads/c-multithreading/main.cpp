#include <iostream>
#include <cstdlib>
#include <pthread.h>
#include <unistd.h>
using namespace std;
#define NUM_THREADS 10

// g++ -o exec main.cpp
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
      cout << "main() : creating thread: " << i << endl;
      rc = pthread_create(&threads[i], NULL, PrintHello, (void *) (size_t) i);
      if (rc) {
         printf("Error: unable to create thread: %d\n", rc);
         exit(-1);
      }
   }

   pthread_exit(NULL);
}
