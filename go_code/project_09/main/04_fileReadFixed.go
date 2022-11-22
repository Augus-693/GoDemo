package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
 * @Project GoDemo
 * @File    04_fileReadFixed.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:21
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	//使用 os 库
	// 创建句柄
	fi, err := os.Open("../file/test.txt")
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)

	// 每次读取 1024 个字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	/*
		//使用 syscall 库
		fd, err := syscall.Open("../file/test.txt", syscall.O_RDONLY, 0)
		if err != nil {
			fmt.Println("Failed on open: ", err)
		}
		defer syscall.Close(fd)

		var wg sync.WaitGroup
		wg.Add(2)
		dataChan := make(chan []byte)
		go func() {
			wg.Done()
			for {
				data := make([]byte, 100)
				n, _ := syscall.Read(fd, data)
				if n == 0 {
					break
				}
				dataChan <- data
			}
			close(dataChan)
		}()

		go func() {
			defer wg.Done()
			for {
				select {
				case data, ok := <-dataChan:
					if !ok {
						return
					}

					fmt.Printf(string(data))
				default:

				}
			}
		}()
		wg.Wait()
	*/
}
