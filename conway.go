package main

import ("fmt"
        //"bufio"
        //"io"
        "io/ioutil"
        "os/exec"
        "os"
        "time"
)

func check(e error) {
    if e != nil{
        panic(e)
    }
}
func clear(){
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}


func parseBoard(data []byte) (stor [][]int){
    var temp []int
    for _, v := range string(data){
        if v == '0'{
            temp = append(temp, 0)
        } else if v ==  '1'{
            temp = append(temp, 1)
        } else if v == '\n'{
            stor = append(stor, temp)
            temp = nil
        }
    }
    return
}


func addAround(x, y int, board [][]int) ( sum int){
    sum = 0
    val := 0
    if x > 0 {
        if y > 0{
            val =  board[x-1][y-1]
     //       if val == 1 { fmt.Printf("%d", 1)}
            sum += val
        }
        if  y <  len(board[0])-1{
            val = board[x-1][y+1]
     //       if val == 1 { fmt.Printf("%d", 2)}
            sum += val
        }
        val = board[x-1][y]
     //   if val == 1 { fmt.Printf("%d", 3)}
        sum += val
    }
    if y > 0{
        val = board[x][y-1]
    //    if val == 1 { fmt.Printf("%d", 4)}
        sum += val
    }
    if  y <  len(board[0])-1{
        val = board[x][y+1]
    //    if val == 1 { fmt.Printf("%d", 5)}
        sum += val
    }
    if x < len(board)-1 {
        if y > 0{
            val = board[x+1][y-1]
    //        if val == 1 { fmt.Printf("%d", 6)}
            sum += val
        }
        if  y <  len(board[0])-1{
            val = board[x+1][y+1]
      //      if val == 1 { fmt.Printf("%d", 7)}
            sum += val
        }
        val = board[x+1][y]
     //   if val == 1 { fmt.Printf("%d", 8)}
        sum += val
    }
    //fmt.Printf("%d", sum)
    return
}

func surround(x int, y int, board [][]int) int{
    sum := 0
    if x  == 0 {
        if y == 0 {
            if board[x+1][y+1] == 1{
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }
            if board[x][y+1] == 1 {
                sum = sum + 1
            }
        } else if y == (len(board[x])-1) {
            if board[x+1][y-1] == 1{
                sum = sum + 1
            }
            if board[x][y-1] == 1 {
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }

        } else{
            if board[x][y-1] == 1{
                sum = sum + 1
            }
            if board[x][y+1] == 1{
                sum = sum + 1
            }
            if board[x+1][y-1] == 1{
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }
            if board[x+1][y+1] == 1{
                sum = sum + 1
            }

        }

    } else if x == (len(board)-1){
        if y == 0 {
            if board[x-1][y+1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1{
                sum = sum +1
            }
            if board[x][y+1] == 1 {
                sum = sum + 1
            }
        } else if y == (len(board[x])-1) {
            if board[x-1][y-1] == 1{
                sum = sum + 1
            }
            if board[x][y-1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1 {
                sum = sum + 1
            }
        } else {
            if board[x][y-1] == 1{
                sum = sum + 1
            }
            if board[x][y+1] == 1{
                sum = sum + 1
            }
            if board[x-1][y-1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1{
                sum = sum + 1
            }
            if board[x-1][y+1] == 1{
                sum = sum + 1
            }
        }
    } else {
        if y == 0 {
            if board[x][y+1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1{
                sum = sum + 1
            }
            if board[x-1][y+1] == 1{
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }
            if board[x+1][y+1] == 1{
                sum = sum + 1
            }
        } else if y == (len(board[x])-1) {
            if board[x][y-1] == 1{
                sum = sum + 1
            }
            if board[x-1][y-1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1{
                sum = sum + 1
            }
            if board[x+1][y-1] == 1{
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }
        } else{
            if board[x][y-1] == 1{
                sum = sum + 1
            }
            if board[x][y+1] == 1{
                sum = sum + 1
            }
            if board[x-1][y-1] == 1{
                sum = sum + 1
            }
            if board[x-1][y] == 1{
                sum = sum + 1
            }
            if board[x-1][y+1] == 1{
                sum = sum + 1
            }
            if board[x+1][y-1] == 1{
                sum = sum + 1
            }
            if board[x+1][y] == 1{
                sum = sum + 1
            }
            if board[x+1][y+1] == 1{
                sum = sum + 1
            }
        }
    }
    fmt.Printf("%d", sum)
    return sum
}


func todo(pos, sum int)(result bool){
    result = false
    if pos == 1{
        if sum < 2 {
            result = false
        } else if sum <= 3{
            result = true
        } else if sum > 3 {
            result = false
        }
    } else if pos == 0{
        if sum == 3{
            result = true
        } else {
            result = false
        }
    }
    return

}




func disp(board [][]int){
    for _, v := range board{
        for _, b := range v{
            if b == 1{
                fmt.Printf("0")
            }else{
                fmt.Printf("_")
            }
        }
        fmt.Printf("\n")
    }

}


func main(){
    dat, err := ioutil.ReadFile("./board")
    var temp int = 0
    check(err)
    clear()
    board := parseBoard(dat)
    disp(board)
    time.Sleep(1000000000)
    clear()
    //addAround(0, 1, board)
    var tempboard [][]int
    tempboard = make( [][]int, len(board))
    for x, _ := range tempboard{
        tempboard[x] = make([]int, len(board[0]))
    }
    for{
        for i, v := range board{
            for j, b := range v{
                temp = addAround(i, j, board)
                if todo(b, temp){
                    tempboard[i][j] = 1
                } else{
                    tempboard[i][j] = 0
                }
            }
      //     fmt.Printf("\n")
       }
       for x, g := range tempboard{
           for y, h := range g{
                temp = h
                board[x][y]=temp
           }
       }
       disp(board)
       time.Sleep(1000000000)
       //fmt.Printf("-----\n")
       clear()
    }
       time.Sleep(500000000)

}
