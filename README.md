# Semaphore Example

Implementation of strong implementation.
```
Semaphore s {
    int count;
    Proc[] queue;
    
    void wait() {
        count--;
        
        if (count < 0) {
            process enqueue to semaphore block queue;
            transit process state to BLOCK;
        }   
    }
    
    void signal() {
        count++;
        
        if (count <= 0) {
            dequeue from block queue;
            transit process state to Runnable;
        }
    }
}
```