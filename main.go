package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readConfigFile(filePath string) (map[string]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    config := make(map[string]string)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "CONFIG_") {
            parts := strings.SplitN(line, "=", 2)
            if len(parts) == 2 {
                key := parts[0]
                value := parts[1]
                config[key] = value
            }
        }
    }

    return config, scanner.Err()
}

func compareConfigFiles(file1Path, file2Path string) {
    config1, err := readConfigFile(file1Path)
    if err != nil {
        fmt.Printf("Cannot read file %s: %v\n", file1Path, err)
        return
    }

    config2, err := readConfigFile(file2Path)
    if err != nil {
        fmt.Printf("Cannot read file %s: %v\n", file2Path, err)
        return
    }

    for key, value1 := range config1 {
        if value2, exists := config2[key]; exists {
            if value1 != value2 {
                fmt.Printf("%s: \x1b[31m%s\x1b[0m != \x1b[32m%s\x1b[0m\n", key, value1, value2)
            }
        } else {
            //fmt.Printf("\x1b[31m%s just in 1: %s\x1b[0m\n", key, value1)
        }
    }
    /*
        for key, value2 := range config2 {
            if _, exists := config1[key]; !exists {
                //fmt.Printf("\x1b[32m%s just in 2: %s\x1b[0m\n", key, value2)
            }
        }
    */
}

func main() {
    // 获取命令行参数
    args := os.Args

    // 检查是否提供了足够的参数
    if len(args) != 3 {
        fmt.Println("Usage: config-compare <file1> <file2>")
        return
    }

    file1Path := args[1]
    file2Path := args[2]

    compareConfigFiles(file1Path, file2Path)
}
