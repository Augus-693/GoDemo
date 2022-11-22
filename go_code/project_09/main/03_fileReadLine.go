package main

/**
 * @Project GoDemo
 * @File    03_fileReadLine.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:18
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	/*
		//bufio.ReadBytes
		// 创建句柄
		fi, err := os.Open("../file/test.txt")
		if err != nil {
			panic(err)
		}

		// 创建 Reader
		r := bufio.NewReader(fi)

		for {
			lineBytes, err := r.ReadBytes('\n')
			line := strings.TrimSpace(string(lineBytes))
			if err != nil && err != io.EOF {
				panic(err)
			}
			if err == io.EOF {
				break
			}
			fmt.Println(line)
		}
	*/

	/*
		//bufio.ReadString
		// 创建句柄
		fi, err := os.Open("../file/test.txt")
		if err != nil {
			panic(err)
		}

		// 创建 Reader
		r := bufio.NewReader(fi)

		for {
			line, err := r.ReadString('\n')
			line = strings.TrimSpace(line)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if err == io.EOF {
				break
			}
			fmt.Println(line)
		}
	*/

	//
}
