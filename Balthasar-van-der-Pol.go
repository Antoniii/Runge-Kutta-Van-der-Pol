// z'' -mu*(1-z^2)*z' + z = 0

package main 

import (
		"fmt"
		"math"
		"os"
		"os/exec" // for exec.Command : Python van Go
		)

func f1(x float64, y float64, z float64) float64 { // y' = mu*(1-z^2)*y - z
	mu := 5.0 // коэффициент, характеризующий нелинейность и силу затухания колебаний 
	return mu*(1-math.Pow(z,2))*y - z
}

func f2(x float64, y float64, z float64) float64 { // z' = y
	return y
}

func main() {
	x0 := 0.0
	y0 := 0.0 // z'(0) = y(0) = 0
	z0 := 2.0 // z(0) = 1
	n := 500 // число шагов
	h := 0.1 // размер шага
	var k1, k2, k3, k4, L1, L2, L3, L4, x, y, z float64
	
	//data := make([][]float64, n)


	file, err := os.Create("data.txt")
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close()


	for i := 1; i <= n; i++  {
		k1 = h*f1(x0,y0,z0)
		L1 = h*f2(x0,y0,z0)

		k2 = h*f1(x0 + h/2, y0 + k1/2, z0 + L1/2)
		L2 = h*f2(x0 + h/2, y0 + k1/2, z0 + L1/2)

		k3 = h*f1(x0 + h/2, y0 + k2/2, z0 + L2/2)
		L3 = h*f2(x0 + h/2, y0 + k2/2, z0 + L2/2)

		k4 = h*f1(x0 + h, y0 + k3, z0 + L3)
		L4 = h*f2(x0 + h, y0 + k3, z0 + L3)

		y = y0 + (k1 + 2*k2 + 2*k3 + k4)/6
		z = z0 + (L1 + 2*L2 + 2*L3 + L4)/6

		x = x0 + h
		fmt.Printf("%.1f %.4f %.4f\n", x, y, z)
		//data[0] = append(data[0], x)
		//data[1] = append(data[1], y)

		_, err = file.WriteString(fmt.Sprintf("%.1f %.4f\n", x, z)) // writing...
        if err != nil {
            fmt.Printf("error writing string: %v", err)
        }

		x0 = x
		y0 = y
		z0 = z
	} 

	//fmt.Printf("%.1f %.4f\n", data[0], data[1])

	
	script := "plot-for-van-der-Pol.py"
	cmd := exec.Command("python", script)
	out, err := cmd.Output()
	fmt.Println(string(out))
	if err != nil {
	    fmt.Println("Error: ", err)
	} 
	
}