package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"
import "fmt"

const grid_size = 3

func main () {
     var factor float32
     var turn bool
     var symbol byte
     var symbol_color color.RGBA
     var bg color.RGBA
     factor = 100
     grid := make ([][]byte, grid_size)
     for i := range grid {
          grid[i] = make([]byte, grid_size)
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
     symbol = ' '
     fmt.Printf ("%s", string(symbol))
     bg = rl.Black
     sel_color := rl.Gray
     sel_color.A = 100
     for !rl.WindowShouldClose() {
          if turn {
               symbol = 'x'
               symbol_color = rl.Red
          } else {
               symbol = 'o'
               symbol_color = rl.Blue
          }
          rl.BeginDrawing()
               rl.ClearBackground (bg)
               x = float32(rl.GetScreenWidth())
               y = float32(rl.GetScreenHeight())
               draw_grid(x, y, rl.Green, y/100)
               mark_grid(x, y, symbol, symbol_color, &grid, sel_color, grid_size)
               if rl.IsKeyPressed (rl.KeyQ) {
                    break
               }
          rl.EndDrawing()
          turn = !turn
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

func mark_grid (x, y float32, symbol byte, symbol_color color.RGBA, grid *[][]byte, grid_bg color.RGBA, grid_size int) {
     sel_color := grid_bg
     sel_color.A = 100
     rec := make ([][]rl.Rectangle, grid_size)
     mouse := rl.GetMousePosition()
     for i := range rec {
          rec[i] = make ([]rl.Rectangle, grid_size)
     }
     rec[0][0] = rl.Rectangle {
          0,
          0,
          x/3,
          y/3,
     }
     rec[0][1] = rl.Rectangle {
          x/3,
          0,
          x/3,
          y/3,
     }
     rec[0][2] = rl.Rectangle {
          2*x/3,
          0,
          x/3,
          y/3,
     }
     rec[1][0] = rl.Rectangle {
          0,
          y/3,
          x/3,
          y/3,
     }
     rec[1][1] = rl.Rectangle {
          x/3,
          y/3,
          x/3,
          y/3,
     }
     rec[1][2] = rl.Rectangle {
          2*x/3,
          y/3,
          x/3,
          y/3,
     }
     rec[2][0] = rl.Rectangle {
          0,
          2*y/3,
          x/3,
          y/3,
     }
     rec[2][1] = rl.Rectangle {
          x/3,
          2*y/3,
          x/3,
          y/3,
     }
     rec[2][2] = rl.Rectangle {
          2*x/3,
          2*y/3,
          x,
          y,
     }
     for i := 0; i < grid_size; i++ {
          for j := 0; j < grid_size; j++ {
               if rl.CheckCollisionPointRec (mouse, rec[i][j]) {
                    rl.DrawRectangleRec (rec[i][j], sel_color)
                    if rl.IsMouseButtonPressed (rl.MouseButtonLeft) {
                         (*grid)[i][j] = symbol
                         font_size := int32(x/10)
                         rl.DrawText (string(symbol), int32(rec[i][j].Width/2)-rl.MeasureText(string(symbol), font_size), int32(rec[i][j].Height/2), font_size, symbol_color)
                    }
               }
          }
     }
}
