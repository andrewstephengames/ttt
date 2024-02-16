package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"

func main () {
     var factor float32
     factor = 100
     rl.InitWindow (int32(16*factor), int32(9*factor), "Tic-Tac-Toe")
     x := 16*factor ; y := 9*factor
     defer rl.CloseWindow()
     for !rl.WindowShouldClose() {
          rl.BeginDrawing()
               rl.ClearBackground (rl.Black)
               draw_grid(x, y, rl.Green, y/100)
               //pos := rl.GetMousePosition()
               //mark_grid (pos, red, blue)
               if rl.IsKeyPressed (rl.KeyQ) {
                    break
               }
          rl.EndDrawing()
     }
}

func draw_grid (x float32, y float32, c color.RGBA, t float32) {
     rl.DrawLineEx(rl.Vector2{0, 0}, rl.Vector2{0, y}, t, c)
     rl.DrawLineEx(rl.Vector2{x/3, 0}, rl.Vector2{x/3, y}, t, c)
     rl.DrawLineEx(rl.Vector2{2*x/3, 0}, rl.Vector2{2*x/3, y}, t, c)
     rl.DrawLineEx(rl.Vector2{x, 0}, rl.Vector2{x, y}, t, c)
     rl.DrawLineEx(rl.Vector2{0, 0}, rl.Vector2{x, 0}, t, c)
     rl.DrawLineEx(rl.Vector2{0, y/3}, rl.Vector2{x, y/3}, t, c)
     rl.DrawLineEx(rl.Vector2{0, 2*y/3}, rl.Vector2{x, 2*y/3}, t, c)
     rl.DrawLineEx(rl.Vector2{0, y}, rl.Vector2{x, y}, t, c)
}

/*
func mark_grid (pos Vector2, x_color color.RGBA, o_color color.RGBA) {
     if pos.x > 0 && pos.x < x/3 && pos.y == 0 {
     }
}
*/
