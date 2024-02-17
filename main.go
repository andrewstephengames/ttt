/*
     TODO: make a simple menu
     TODO: generalize on grid_size > 3
*/

package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"
import "fmt"

const grid_size = 3

func main () {
     var factor float32
     var symbol_color color.RGBA
     var bg color.RGBA
     factor = 100
     grid := make ([][]byte, grid_size)
     for i := range grid {
          grid[i] = make([]byte, grid_size)
     }
     rec := make ([][]rl.Rectangle, grid_size)
     for i := range rec {
          rec[i] = make ([]rl.Rectangle, grid_size)
     }
     for i := 0; i < grid_size; i++ {
          for j := 0; j < grid_size; j++ {
               grid[i][j] = ' '
          }
     }
     rl.SetTraceLogLevel (rl.LogError)
     rl.InitWindow (int32(16*factor), int32(9*factor), "Tic-Tac-Toe")
     x := 16*factor ; y := 9*factor
     rl.SetConfigFlags (rl.FlagWindowResizable)
     rl.SetTargetFPS(60)
     defer rl.CloseWindow()
     bg = rl.Black
     sel_color := rl.Gray
     sel_color.A = 100
     for !rl.WindowShouldClose() {
          rl.BeginDrawing()
               rl.ClearBackground (bg)
               x = float32(rl.GetScreenWidth())
               y = float32(rl.GetScreenHeight())
               font_size := int32(x/10)
               draw_grid(x, y, rl.Green, y/100)
               mark_grid(x, y, &grid, &rec, sel_color, grid_size)
               for i := 0; i < grid_size; i++ {
                    for j := 0; j < grid_size; j++ {
                         textX := int32(rec[i][j].X + rec[i][j].Width/2) - rl.MeasureText(string(grid[i][j]), font_size)/2
                         textY := int32(rec[i][j].Y + rec[i][j].Height/2) - font_size/2
                         if grid[i][j] == 'x' {
                              symbol_color = rl.Red
                         } else {
                              symbol_color = rl.Blue
                         }
                         rl.DrawText(string(grid[i][j]), textX, textY, font_size, symbol_color)
                         if check_condition (&grid, grid_size) == 'x' {
                              msg := "x won"
                              rl.DrawText (msg, int32(x/2)-rl.MeasureText (msg, int32(factor))/2, int32(y/2)-int32(factor/2), int32(factor), rl.Yellow)
                         } else if check_condition (&grid, grid_size) == 'o' {
                              msg := "o won"
                              rl.DrawText (msg, int32(x/2)-rl.MeasureText (msg, int32(factor))/2, int32(y/2)-int32(factor/2), int32(factor), rl.Yellow)
                         }
                    }
               }
               if rl.IsKeyPressed (rl.KeyQ) {
                    break
               }
               if rl.IsKeyPressed (rl.KeyR) {
                    reset_grid (&grid, grid_size)
               }
          rl.EndDrawing()
     }
     test_grid(grid, grid_size)
}

func draw_grid (x, y float32, c color.RGBA, t float32) {
     rl.DrawLineEx(rl.Vector2{0, 0}, rl.Vector2{0, y}, t, c)
     rl.DrawLineEx(rl.Vector2{x/3, 0}, rl.Vector2{x/3, y}, t, c)
     rl.DrawLineEx(rl.Vector2{2*x/3, 0}, rl.Vector2{2*x/3, y}, t, c)
     rl.DrawLineEx(rl.Vector2{x, 0}, rl.Vector2{x, y}, t, c)
     rl.DrawLineEx(rl.Vector2{0, 0}, rl.Vector2{x, 0}, t, c)
     rl.DrawLineEx(rl.Vector2{0, y/3}, rl.Vector2{x, y/3}, t, c)
     rl.DrawLineEx(rl.Vector2{0, 2*y/3}, rl.Vector2{x, 2*y/3}, t, c)
     rl.DrawLineEx(rl.Vector2{0, y}, rl.Vector2{x, y}, t, c)
}

func test_grid (grid [][]byte, grid_size int) {
     for i := 0; i < grid_size; i++ {
          for j := 0; j < grid_size; j++ {
               fmt.Print (string(grid[i][j]), " ")
          }
          fmt.Println()
     }
}

func mark_grid (x, y float32, grid *[][]byte, rec *[][] rl.Rectangle, grid_bg color.RGBA, grid_size int) {
     sel_color := grid_bg
     sel_color.A = 100
     mouse := rl.GetMousePosition()
     (*rec)[0][0] = rl.Rectangle {
          0,
          0,
          x/3,
          y/3,
     }
     (*rec)[0][1] = rl.Rectangle {
          x/3,
          0,
          x/3,
          y/3,
     }
     (*rec)[0][2] = rl.Rectangle {
          2*x/3,
          0,
          x/3,
          y/3,
     }
     (*rec)[1][0] = rl.Rectangle {
          0,
          y/3,
          x/3,
          y/3,
     }
     (*rec)[1][1] = rl.Rectangle {
          x/3,
          y/3,
          x/3,
          y/3,
     }
     (*rec)[1][2] = rl.Rectangle {
          2*x/3,
          y/3,
          x/3,
          y/3,
     }
     (*rec)[2][0] = rl.Rectangle {
          0,
          2*y/3,
          x/3,
          y/3,
     }
     (*rec)[2][1] = rl.Rectangle {
          x/3,
          2*y/3,
          x/3,
          y/3,
     }
     (*rec)[2][2] = rl.Rectangle {
          2*x/3,
          2*y/3,
          x/3,
          y/3,
     }
     for i := 0; i < grid_size; i++ {
          for j := 0; j < grid_size; j++ {
               if rl.CheckCollisionPointRec (mouse, (*rec)[i][j]) {
                    rl.DrawRectangleRec ((*rec)[i][j], sel_color)
                    if (*grid)[i][j] == ' ' {
                         if rl.IsMouseButtonPressed (rl.MouseButtonLeft) {
                              (*grid)[i][j] = 'x'
                         }
                         if rl.IsMouseButtonPressed (rl.MouseButtonRight) {
                              (*grid)[i][j] = 'o'
                         }
                    }
               }
          }
     }
}

func check_condition (grid *[][]byte, grid_size int32) byte {
     var i int32
     var j int32
     j = 1
     for i = 0; i < grid_size; i += 2 {
          if (*grid)[i][j] == (*grid)[i][j-1] && (*grid)[i][j] == (*grid)[i][j+1] &&
             (*grid)[i][j] != ' ' {
               return (*grid)[i][j]
          }
     }
     i = 1 ; j = 1
     if (*grid)[i][j] == (*grid)[i-1][j-1] && (*grid)[i][j] == (*grid)[i+1][j+1] &&
        (*grid)[i][j] != ' ' {
          return (*grid)[i][j]
     }
     if (*grid)[i][j] == (*grid)[i-1][j+1] && (*grid)[i][j] == (*grid)[i+1][j-1] {
          return (*grid)[i][j]
     }
     for j = 0; j < grid_size; j++ {
          if (*grid)[i][j] == (*grid)[i-1][j] && (*grid)[i][j] == (*grid)[i+1][j] &&
             (*grid)[i][j] != ' ' {
               return (*grid)[i][j]
          }
     }
     return ' '
}

func reset_grid (grid *[][]byte, grid_size int32) {
     var i int32
     var j int32
     for i = 0; i < grid_size; i++ {
          for j = 0; j < grid_size; j++ {
               (*grid)[i][j] = ' '
          }
     }
}
