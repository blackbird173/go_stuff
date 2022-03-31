package main
import "mocha"
func main() {
	ring := "[-------------]\n"
	mocha.Rangfunc(0, 4, func() {
		ring += "[.............]\n"
	})
	ring += "[-------------]"
	inp := ""
	pos := 18
	pos2 := 76
	uhealth := 3
	ehealth := 3
	ring = mocha.Replace(ring, pos, '|')
	ring = mocha.Replace(ring, pos2, '}')
	for {
		if ehealth == 0 {
			mocha.Appendfile("boxinglogs.txt", "u win\n")
			break
		} else if uhealth == 0 {
			mocha.Appendfile("boxinglogs.txt", "u lose\n")
			break
		}
		go func() {
			inp = mocha.Input(inp)
		}()
		mocha.Sleep(1000)
		switch inp {
		case "w":
			if ring[pos-16:pos-15] == "." {
				pos -= 16
				ring = mocha.Replace(ring, pos, '|')
				ring = mocha.Replace(ring, pos+16, '.')
				inp = ""
			}
		case "s":
			if ring[pos+16:pos+17] == "." {
				pos += 16
				ring = mocha.Replace(ring, pos, '|')
				ring = mocha.Replace(ring, pos-16, '.')
				inp = ""
			}
		case "a":
			if ring[pos-1:pos] == "." {
				pos--
				ring = mocha.Replace(ring, pos, '|')
				ring = mocha.Replace(ring, pos+1, '.')
				inp = ""
			}
		case "d":
			if ring[pos+1:pos+2] == "." {
				pos++
				ring = mocha.Replace(ring, pos, '|')
				ring = mocha.Replace(ring, pos-1, '.')
				inp = ""
			}
		case "f":
			if ring[pos+1:pos+2] == "}" || ring[pos-1:pos] == "}" || ring[pos-16:pos-15] == "}" || ring[pos+16:pos+17] == "}" {
				ehealth--
				println("he lost 1 point")
			}
		}
		num := mocha.Random(4)
		switch num {
		case 0:
			if ring[pos2-16:pos2-15] == "." {
				pos2 -= 16
				ring = mocha.Replace(ring, pos2, '}')
				ring = mocha.Replace(ring, pos2+16, '.')
			}
		case 1:
			if ring[pos2+16:pos2+17] == "." {
				pos2 += 16
				ring = mocha.Replace(ring, pos2, '}')
				ring = mocha.Replace(ring, pos2-16, '.')
			}
		case 2:
			if ring[pos2-1:pos2] == "." {
				pos2--
				ring = mocha.Replace(ring, pos2, '}')
				ring = mocha.Replace(ring, pos2+1, '.')
			}
		case 3:
			if ring[pos2+1:pos2+2] == "." {
				pos2++
				ring = mocha.Replace(ring, pos2, '}')
				ring = mocha.Replace(ring, pos2-1, '.')
			}
			
		case 4:
			if ring[pos2-1:pos2] == "|" || ring[pos2+1:pos2+2] == "|" || ring[pos2-16:pos2-15] == "|" || ring[pos2+16:pos2+17] == "|" {
				uhealth--
				println("u lost 1 point")
			}
		}
		println(ring)
	}
}
