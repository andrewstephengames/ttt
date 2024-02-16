package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"

func main () {
     black := color.RGBA {0, 0, 0, 255}
     //red := color.RGBA {255, 0, 0, 255}
     green := color.RGBA {0, 255, 0, 255}
     var factor int32
     factor = 100
     rl.InitWindow (16*factor, 9*factor, "Tic-Tac-Toe")
     x := 16*factor ; y := 9*factor
     defer rl.CloseWindow()
     for !rl.WindowShouldClose() {
          rl.BeginDrawing()
               rl.ClearBackground (black)
               rl.DrawLine(0, 0, 0, y, green)
               rl.DrawLine(x/3, 0, x/3, y, green)
               rl.DrawLine(2*x/3, 0, 2*x/3, y, green)
               rl.DrawLine(x, 0, x, y, green)
               rl.DrawLine(0, 0, x, 0, green)
               rl.DrawLine(0, y/3, x, y/3, green)
               rl.DrawLine(0, 2*y/3, x, 2*y/3, green)
               rl.DrawLine(0, y, x, y, green)
          rl.EndDrawing()
     }
}
