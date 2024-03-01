/*
     TODO: make a simple menu
     TODO: generalize on grid_size > 3
*/

package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"
import "fmt"
import "os"
import "log"
import "math/rand"

const grid_size = 3
const (
     Menu int = iota
     Game
     End
)

const (
     play_button int = iota
     quit_button
     last
)

type Button struct {
     label string
     box rl.Rectangle
     bg color.RGBA
     fg color.RGBA
}

func main () {
     var factor float32
     var symbol byte
     var bg color.RGBA
     buttons := make ([]Button, last)
     button_color := rl.SkyBlue
     init_button (&buttons[play_button], button_color, rl.White, "Play")
     init_button (&buttons[quit_button], button_color, rl.White, "Quit")
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
     rl.SetConfigFlags (rl.FlagWindowResizable)
     rl.InitWindow (int32(16*factor), int32(9*factor), "Tic-Tac-Toe")
     x := 16*factor ; y := 9*factor
     rl.SetTargetFPS(60)
     defer rl.CloseWindow()
     bg = rl.Black
     sel_color := rl.Gray
     sel_color.A = 100
     turn := rand.Intn(2) == 1
     init_state := Menu
     set_turn (&turn, &symbol)
     for !rl.WindowShouldClose() {
          rl.BeginDrawing()
               rl.ClearBackground (bg)
               x = float32(rl.GetScreenWidth())
               y = float32(rl.GetScreenHeight())
               draw_grid(x, y, rl.SkyBlue, y/100)
               mark_grid(x, y, &grid, &rec, sel_color, grid_size, &turn, &symbol)
               for i := 0; i < grid_size; i++ {
                    for j := 0; j < grid_size; j++ {
                         state_machine (&init_state, x, y, &buttons, &rec, &grid, grid_size, factor);
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

func init_button (button *Button, bg color.RGBA, fg color.RGBA, label string) {
     (*button).bg = bg
     (*button).fg = fg
     (*button).label = label
}

func get_sel_color () color.RGBA {
     var scale uint8
     scale = 20
     sel_color := color.RGBA { rl.SkyBlue.R - scale, rl.SkyBlue.G - scale, rl.SkyBlue.B - scale, uint8(255)}
     return sel_color
}

func state_machine (game_state *int, x float32, y float32, button *[]Button, rec *[][]rl.Rectangle, grid *[][]byte, grid_size int, factor float32) {
     var canvas rl.Vector2
     var canvas2 rl.Vector2
     var font_size int32
     var game_font_size int32
     var symbol_color color.RGBA
     canvas.X = x;
     canvas.Y = y;
     switch (*game_state) {
          case Menu:
               rl.DrawRectangle (0, 0, int32(canvas.X), int32(canvas.Y), rl.SkyBlue);
               title := "Tic-Tac-Toe"
               font_size = int32(canvas.Y/10)
               canvas2.X = canvas.X/2 - float32(rl.MeasureText (title, font_size)/2)
               canvas2.Y = canvas.Y/4 - float32(font_size/2)
               rl.DrawText (title, int32(canvas2.X), int32(canvas2.Y), font_size, rl.White)

               text_x := canvas.X/2 - float32(rl.MeasureText ((*button)[play_button].label, font_size)/2)
               text_y := canvas.Y/2
               text_box_w := float32(rl.MeasureText ((*button)[play_button].label, font_size))
               text_box_h := float32(font_size)

               (*button)[play_button].box.Width = text_box_w
               (*button)[play_button].box.Height = text_box_h
               (*button)[play_button].box.X = text_x
               (*button)[play_button].box.Y = text_y
               if rl.CheckCollisionPointRec (rl.GetMousePosition(), (*button)[play_button].box) {
                    (*button)[play_button].bg = get_sel_color()
                    if rl.IsMouseButtonPressed (rl.MouseButtonLeft) {
                         *game_state = Game
                    }
               } else {
                    (*button)[play_button].bg = rl.SkyBlue
               }
               rl.DrawRectangleRec ((*button)[play_button].box, (*button)[play_button].bg)
               rl.DrawText ((*button)[play_button].label, int32(text_x), int32(text_y), font_size, (*button)[play_button].fg)

               text_x = canvas.X/2 - float32(rl.MeasureText ((*button)[quit_button].label, font_size)/2)
               text_y = 3*canvas.Y/4
               text_box_w = float32(rl.MeasureText ((*button)[quit_button].label, font_size))
               text_box_h = float32(font_size)

               (*button)[quit_button].box.Width = text_box_w
               (*button)[quit_button].box.Height = text_box_h
               (*button)[quit_button].box.X = text_x
               (*button)[quit_button].box.Y = text_y
               if rl.CheckCollisionPointRec (rl.GetMousePosition(), (*button)[quit_button].box) {
                    (*button)[quit_button].bg = get_sel_color()
                    if rl.IsMouseButtonPressed (rl.MouseButtonLeft) {
                         *game_state = End
                    }
               } else {
                    (*button)[quit_button].bg = rl.SkyBlue
               }
               rl.DrawRectangleRec ((*button)[quit_button].box, (*button)[quit_button].bg)
               rl.DrawText ((*button)[quit_button].label, int32(text_x), int32(text_y), font_size, (*button)[quit_button].fg)

               if rl.IsKeyPressed (rl.KeyQ) || rl.IsKeyPressed (rl.KeyEscape) {
                    *game_state = End
               }
               break
          case Game:
               game_font_size = int32(x/10)
               for i := 0; i < grid_size; i++ {
                    for j := 0; j < grid_size; j++ {
                         textX := int32((*rec)[i][j].X + (*rec)[i][j].Width/2) - rl.MeasureText(string((*grid)[i][j]), game_font_size)/2
                         textY := int32((*rec)[i][j].Y + (*rec)[i][j].Height/2) - game_font_size/2
                         if (*grid)[i][j] == 'x' {
                              symbol_color = rl.Red
                         } else {
                              symbol_color = rl.Blue
                         }
                         rl.DrawText(string((*grid)[i][j]), textX, textY, game_font_size, symbol_color)
                         if check_condition (grid, int32(grid_size)) == 'x' {
                              msg := "x won"
                              rl.DrawText (msg, int32(x/2)-rl.MeasureText (msg, int32(factor))/2, int32(y/2)-int32(factor/2), int32(factor), rl.Yellow)
                         } else if check_condition (grid, int32(grid_size)) == 'o' {
                              msg := "o won"
                              rl.DrawText (msg, int32(x/2)-rl.MeasureText (msg, int32(factor))/2, int32(y/2)-int32(factor/2), int32(factor), rl.Yellow)
                         } else if is_draw (grid, int32(grid_size)) {
                              msg := "draw"
                              rl.DrawText (msg, int32(x/2)-rl.MeasureText (msg, int32(factor))/2, int32(y/2)-int32(factor/2), int32(factor), rl.Yellow)
                         }
                    }
               }
               if rl.IsKeyPressed (rl.KeyQ) || rl.IsKeyPressed (rl.KeyEscape) {
                    *game_state = Menu
               }
               break
          case End:
               os.Exit(0);
               break
          default: {
               log.Fatal ("Unknown game state!")
          }
     }
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

func set_turn (turn *bool, symbol *byte) {
     if *turn {
          *symbol = 'x'
     } else {
          *symbol = 'o'
     }
     *turn = !*turn
}

func mark_grid (x, y float32, grid *[][]byte, rec *[][] rl.Rectangle, grid_bg color.RGBA, grid_size int, turn *bool, symbol *byte) {
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
                              set_turn (turn, symbol)
                              (*grid)[i][j] = *symbol
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
     if (*grid)[i][j] == (*grid)[i-1][j+1] && (*grid)[i][j] == (*grid)[i+1][j-1] &&
        (*grid)[i][j] != ' ' {
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

func is_draw (grid *[][]byte, grid_size int32) bool {
     var i, j int32
     has_space := false
     for i = 0; i < grid_size; i++ {
          for j = 0; j < grid_size; j++ {
               if (*grid)[i][j] == ' ' {
                    has_space = true
               }
          }
     }
     return !has_space
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
