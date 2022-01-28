#include <pthread.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdatomic.h>

#define PING "ping"
#define PONG "pong"

void critical(const char * str) {
  size_t len = strlen(str);
  for (size_t i = 0; i < len; ++i) {
    printf("%c", str[i]);
  } // for
  printf("\n");
} // str

typedef atomic_flag mut_t;
volatile mut_t mut = ATOMIC_FLAG_INIT; // false; true = locked; false = unlocked

// void acquire(mut_t * m)
#define acquire(m) while (atomic_flag_test_and_set(m))
// void release(mut_t * m)
#define release(m) atomic_flag_clear(m)

void * pingpong(void * p) {
  char * msg = (char *) p;
  for (;;) {
    acquire(&mut);
    critical(msg);
    release(&mut);
  } // for
} // pingpong

// gcc main.c -o exec
// ./exec
int main() {
  setvbuf(stdout, NULL, _IONBF, 0);
  pthread_t ping_thread;
  pthread_t pong_thread;
  pthread_create(&ping_thread, NULL, pingpong, PING);
  pthread_create(&pong_thread, NULL, pingpong, PONG);
  for(;;);
  return 0;
}
