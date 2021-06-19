#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>
#include <work.h>

void *threadPrint(void *threadid) {
   long tid;
   tid = (long)threadid;
   sleep(1);
   printf("thread %ld: printing\n", tid);
   sleep(3);
   printf("thread %ld: exiting\n", tid);
   pthread_exit(NULL);
}

int work (struct Thread *t) {
   pthread_t thread;
   int rc;
   rc = pthread_create(&thread, NULL, threadPrint, (void *) (size_t) t->id);
   if (rc) {
      printf("Error: unable to create thread: %d\n", rc);
      exit(-1);
   }
   pthread_join(thread, NULL);
   return 0;
}
