package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate processing time
        results <- job * 2 // Process job (e.g., double it)
        fmt.Printf("Worker %d processed job %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    var wg sync.WaitGroup

    // Start workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Close results channel once all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()

    // Read results
    for result := range results {
        fmt.Println("Result:", result)
    }
    fmt.Println("All jobs processed.")
}
