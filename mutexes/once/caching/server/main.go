package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {
	errChan := make(chan error)
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, os.Interrupt)

	go func() {
		fmt.Println("api gateway successfully started")
		err := http.ListenAndServe("localhost:8080", apiGateway())
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		fmt.Println("users api successfully started")
		err := http.ListenAndServe("localhost:8081", usersAPI())
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		fmt.Println("order api successfully started")
		err := http.ListenAndServe("localhost:8082", ordersAPI())
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		fmt.Println("products api successfully started")
		err := http.ListenAndServe("localhost:8083", productsAPI())
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-quitChan:
		fmt.Println("successfully exited")
	case err := <-errChan:
		fmt.Println("could not start server:", err)
	}
}

type analyticsResponse struct {
	User     user      `json:"user"`
	Products []product `json:"products"`
}

func apiGateway() *http.ServeMux {
	userClient := client{
		url:   "localhost:8081/api/users",
		cache: map[string]*cacheEntry{},
		mu:    new(sync.Mutex),
	}
	ordersClient := client{
		url:   "localhost:8082/api/orders",
		cache: map[string]*cacheEntry{},
		mu:    new(sync.Mutex),
	}
	productClient := client{
		url:   "localhost:8083/api/products",
		cache: map[string]*cacheEntry{},
		mu:    new(sync.Mutex),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/gateway/products", func(w http.ResponseWriter, r *http.Request) {
		productID := r.URL.Query().Get("id")
		if strings.TrimSpace(productID) == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		res := productClient.query(productID, "id="+productID)
		w.Header().Set("content-type", "application/json")
		_, err := w.Write(append(res, '\n'))
		if err != nil {
			log.Fatalf("could not write json: %v", err)
		}
	})
	mux.HandleFunc("/api/gateway/analytics", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if strings.TrimSpace(userID) == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		userRes := userClient.query(userID, "id="+userID)
		ordersRes := ordersClient.query(userID, "user_id="+userID)
		fmt.Println("userRes", string(userRes))
		fmt.Println("ordersRes", string(ordersRes))
		res := analyticsResponse{}

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatalf("could not encode json response: %v", err)
		}
	})
	return mux
}

type user struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
}

func usersAPI() *http.ServeMux {
	usersDB := map[string]user{
		"1": {ID: "1", Email: "user1@example.com"},
		"2": {ID: "2", Email: "user2@example.com"},
		"3": {ID: "3", Email: "user3@example.com"},
		"4": {ID: "4", Email: "user4@example.com"},
		"5": {ID: "5", Email: "user5@example.com"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		time.Sleep(time.Second)
		user, found := usersDB[userID]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Fatalf("could not encode user: %v", err)
		}
	})
	return mux
}

type order struct {
	ID         string   `json:"id"`
	UserID     string   `json:"user_id"`
	ProductIDs []string `json:"product_ids"`
}

func ordersAPI() *http.ServeMux {
	ordersDB := map[string][]order{
		"1": {
			{ID: "1", UserID: "1", ProductIDs: []string{"1", "2"}},
			{ID: "2", UserID: "1", ProductIDs: []string{"3"}},
		},
		"2": {
			{ID: "3", UserID: "2", ProductIDs: []string{"1"}},
		},
		"3": {
			{ID: "4", UserID: "3", ProductIDs: []string{"2"}},
		},
		"4": {
			{ID: "5", UserID: "4", ProductIDs: []string{"3"}},
		},
		"5": {
			{ID: "6", UserID: "5", ProductIDs: []string{"1"}},
			{ID: "7", UserID: "5", ProductIDs: []string{"2"}},
			{ID: "8", UserID: "5", ProductIDs: []string{"3"}},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		orders, found := ordersDB[userID]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		time.Sleep(time.Second)
		err := json.NewEncoder(w).Encode(orders)
		if err != nil {
			log.Fatalf("could not encode orders: %v", err)
		}
	})
	return mux
}

type product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func productsAPI() *http.ServeMux {
	productsDB := map[string]product{
		"1": {ID: "1", Name: "Coca Cola - 2L", Price: 1.49},
		"2": {ID: "2", Name: "Red Bull - Original", Price: 3.99},
		"3": {ID: "3", Name: "Starbucks - Iced Coffee", Price: 2.25},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		productID := r.URL.Query().Get("id")
		product, found := productsDB[productID]
		fmt.Println("found", found)
		if !found {
			fmt.Println("here")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		time.Sleep(time.Second)
		err := json.NewEncoder(w).Encode(product)
		if err != nil {
			log.Fatalf("could not encode products: %v", err)
		}
	})
	return mux
}

type cacheEntry struct {
	data []byte
	once *sync.Once
}

type client struct {
	cache map[string]*cacheEntry
	mu    *sync.Mutex
	url   string
}

func (c *client) query(key, query string) []byte {
	c.mu.Lock()
	entry, found := c.cache[key]
	if !found {
		entry = &cacheEntry{once: new(sync.Once)}
		c.cache[key] = entry
	}
	c.mu.Unlock()

	entry.once.Do(func() {
		uri := fmt.Sprintf("http://%s?%s", c.url, url.PathEscape(query))
		res, err := http.Get(uri)
		if err != nil {
			log.Fatalf("could not fetch data from external api %s: %v", c.url, err)
		}

		if res.StatusCode >= 300 {
			c.mu.Lock()
			delete(c.cache, key)
			c.mu.Unlock()
			return
		}

		entry.data, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("could not read response body: %v", err)
		}
	})

	return entry.data
}
