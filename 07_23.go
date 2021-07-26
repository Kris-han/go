package _go


//Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
//而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码


k := 6
switch k {
case 4: fmt.Println("was <= 4"); fallthrough;
case 5: fmt.Println("was <= 5"); fallthrough;
case 6: fmt.Println("was <= 6");
case 7: fmt.Println("was <= 7"); fallthrough;
case 8: fmt.Println("was <= 8"); fallthrough;
default: fmt.Println("default case")
}

// 输出
//was <= 6

k := 6
switch k {
case 4: fmt.Println("was <= 4"); fallthrough;
case 5: fmt.Println("was <= 5"); fallthrough;
case 6: fmt.Println("was <= 6"); fallthrough;
case 7: fmt.Println("was <= 7"); fallthrough;
case 8: fmt.Println("was <= 8"); fallthrough;
default: fmt.Println("default case")
}

// 输出
//was <= 6
//was <= 7
//was <= 8
//default case



