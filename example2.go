var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string

// 2.3.1. Short Variable Declarations
// Wi thin a function, an alternate for m called a sh ort var iab le declarati on may be used to declare
// and initialize local var iables. It takes the for m name := expression, and the typ e of name is
// deter mined by the typ e of expression. Here are three of the many short var iable decl arat ions
// in the lissajous func tion (§1.4):
// anim := gif.GIF{LoopCount: nframes}
// freq := rand.Float64() * 3.0
// t := 0.0