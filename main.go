package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//neighbours < 2 => d
//neighbours == 2 or 3 => l
//neighbours == 3 and you are dead => l

func CallClear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Log().Err(err)
	}
}

func main() {
	//sigc := make(chan os.Signal, 1)
	//signal.Notify(sigc,
	//	syscall.SIGHUP,
	//	syscall.SIGINT,
	//	syscall.SIGTERM,
	//	syscall.SIGQUIT)
	//go func() {
	//	switch <-sigc {
	//	case os.Interrupt:
	//
	//	case syscall.SIGTERM:
	//		//handle SIGTERM
	//	}
	//}()
	num := 0
	size := 16
	fmt.Print("Enter matrix size: ")
	fmt.Fscan(os.Stdin, &size)
	speed := 2
	fmt.Print("\nEnter update speed in seconds: ")
	fmt.Fscan(os.Stdin, &speed)
	CallClear()
	n := make([][]int, size)
	for i := 0; i < size; i++ {
		n[i] = make([]int, size)
	}
	generate(n)
	beg := make([][]int, size)
	for i := 0; i < size; i++ {
		beg[i] = make([]int, size)
		for j := 0; j < size; j++ {
			beg[i][j] = n[i][j]
		}
	}
	for {
		show(n, num, beg)
		gameOfLife(n)
		time.Sleep(time.Duration(speed) * time.Second)
		CallClear()
		num++
	}
}

func generate(b [][]int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = rand.Intn(2)
		}
	}
}

func show(b [][]int, num int, beg [][]int) {
	sizer := len(b)
	sizec := len(b[0])
	for i := 0; i < sizer; i++ {
		for j := 0; j < sizec; j++ {
			fmt.Print(b[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("number of generation: " + strconv.Itoa(num))
	//for i := 0; i < sizer; i++ {
	//	for j := 0; j < sizec; j++ {
	//		fmt.Print(beg[i][j])
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}
}

func gameOfLife(board [][]int)  {
	m, t := len(board), len(board[0])
	n := make([][]int, m)
	for i := range n {
		n[i] = make([]int, t)
		copy(n[i], board[i])
	}
	c := 0
	last := len(board) - 1
	last2 := len(board[0]) - 1
	if last > 0 && last2 > 0{
		if board[0][0] == 0 {
			if board[0][1] == 1 && board[1][0] == 1 && board[1][1] == 1 {
				n[0][0] = 1
			}
		} else {
			if board[0][1] == 1 {
				c+=1
			}
			if board[1][1] == 1 {
				c+=1
			}
			if board[1][0] == 1 {
				c+=1
			}
			if c > 1 {
				n[0][0] = 1
			} else {
				n[0][0] = 0
			}
		}
		c = 0
		if board[0][last2] == 0 {
			if board[0][last2-1] == 1 && board[1][last2-1] == 1 && board[1][last2] == 1 {
				n[0][last2] = 1
			}
		} else {
			if board[0][last2-1] == 1 {
				c+=1
			}
			if board[1][last2] == 1 {
				c+=1
			}
			if board[1][last2-1] == 1 {
				c+=1
			}
			if c > 1 {
				n[0][last2] = 1
			} else {
				n[0][last2] = 0
			}
		}
		c = 0
		if board[last][last2] == 0 {
			if board[last][last2-1] == 1 && board[last-1][last2-1] == 1 && board[last-1][last2] == 1 {
				n[last][last2] = 1
			}
		} else {
			if board[last][last2-1] == 1 {
				c+=1
			}
			if board[last-1][last2] == 1 {
				c+=1
			}
			if board[last-1][last2-1] == 1 {
				c+=1
			}
			if c > 1 {
				n[last][last2] = 1
			} else {
				n[last][last2] = 0
			}
		}
		c = 0
		if board[last][0] == 0 {
			if board[last][1] == 1 && board[last-1][0] == 1 && board[last-1][1] == 1 {
				n[last][0] = 1
			}
		} else {
			if board[last][1] == 1 {
				c+=1
			}
			if board[last-1][1] == 1 {
				c+=1
			}
			if board[last-1][0] == 1 {
				c+=1
			}
			if c > 1 {
				n[last][0] = 1
			} else {
				n[last][0] = 0
			}
		}
		c = 0
		if last > 1 {
			for i := 1; i < last2; i++ {
				if board[0][i-1] == 1 {
					c+=1
				}
				if board[0][i+1] == 1 {
					c+=1
				}
				if board[1][i-1] == 1 {
					c+=1
				}
				if board[1][i] == 1 {
					c+=1
				}
				if board[1][i+1] == 1 {
					c+=1
				}
				if board[0][i] == 0 {
					if c == 3 {
						n[0][i] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[0][i] = 1
					} else {
						n[0][i] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				if board[i-1][0] == 1 {
					c+=1
				}
				if board[i+1][0] == 1 {
					c+=1
				}
				if board[i-1][1] == 1 {
					c+=1
				}
				if board[i][1] == 1 {
					c+=1
				}
				if board[i+1][1] == 1 {
					c+=1
				}
				if board[i][0] == 0 {
					if c == 3 {
						n[i][0] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[i][0] = 1
					} else {
						n[i][0] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				if board[i-1][last2] == 1 {
					c+=1
				}
				if board[i+1][last2] == 1 {
					c+=1
				}
				if board[i-1][last2-1] == 1 {
					c+=1
				}
				if board[i][last2-1] == 1 {
					c+=1
				}
				if board[i+1][last2-1] == 1 {
					c+=1
				}
				if board[i][last2] == 0 {
					if c == 3 {
						n[i][last2] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[i][last2] = 1
					} else {
						n[i][last2] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last2; i++ {
				if board[last][i-1] == 1 {
					c+=1
				}
				if board[last][i+1] == 1 {
					c+=1
				}
				if board[last-1][i-1] == 1 {
					c+=1
				}
				if board[last-1][i] == 1 {
					c+=1
				}
				if board[last-1][i+1] == 1 {
					c+=1
				}
				if board[last][i] == 0 {
					if c == 3 {
						n[last][i] = 1
					}
				} else {
					if c == 3 || c == 2 {
						n[last][i] = 1
					} else {
						n[last][i] = 0
					}
				}
				c = 0
			}
			for i := 1; i < last; i++ {
				for j := 1; j < last2; j++ {
					if board[i-1][j-1] == 1 {
						c+=1
					}
					if board[i-1][j] == 1 {
						c+=1
					}
					if board[i-1][j+1] == 1 {
						c+=1
					}
					if board[i+1][j-1] == 1 {
						c+=1
					}
					if board[i+1][j] == 1 {
						c+=1
					}
					if board[i+1][j+1] == 1 {
						c+=1
					}
					if board[i][j-1] == 1 {
						c+=1
					}
					if board[i][j+1] == 1 {
						c+=1
					}
					if board[i][j] == 0 {
						if c == 3 {
							n[i][j] = 1
						}
					} else {
						if c == 3 || c == 2 {
							n[i][j] = 1
						} else {
							n[i][j] = 0
						}
					}
					c = 0
				}
			}
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else if last2 > 0 {
		n[0][0] = 0
		n[0][last2] = 0
		for i := 1; i < last2; i++ {
			if board[0][i-1] == 1 {
				c+=1
			}
			if board[0][i+1] == 1 {
				c+=1
			}
			if board[0][i] == 1 {
				if c == 2 {
					n[0][i] = 1
				} else {
					n[0][i] = 0
				}
			}
			c = 0
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else if last > 0 {
		n[0][0] = 0
		n[last][0] = 0
		for i := 1; i < last; i++ {
			if board[i-1][0] == 1 {
				c+=1
			}
			if board[i+1][0] == 1 {
				c+=1
			}
			if board[i][0] == 1 {
				if c == 2 {
					n[i][0] = 1
				} else {
					n[i][0] = 0
				}
			}
			c = 0
		}
		for i := range n {
			for j := range n[i] {
				board[i][j] = n[i][j]
			}
		}
	} else {
		board[0][0] = 0
	}
}