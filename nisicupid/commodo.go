package main

import (
    "fmt"
    "errors"
)

// Assume GetInstance is a function that returns an instance of a service
func GetInstance(name, config string, initializer func() (interface{}, error)) (interface{}, error) {
    // Here you might have some logic to fetch or initialize the instance
    instance, err := initializer()
    if err != nil {
        return nil, err
    }
    return instance, nil
}

func main() {
    // Example usage of GetInstance for "usersvc"
    usersvc, err := GetInstance("usersvc", "default", func() (interface{}, error) {
        // Replace with actual initialization logic for "usersvc"
        // For example:
        // return NewUserService(), nil
        return nil, errors.New("usersvc initialization not implemented")
    })

    if err != nil {
        fmt.Println("Error initializing usersvc:", err)
        return
    }

    fmt.Println("usersvc initialized successfully:", usersvc)
}
