package basics;

type Details struct {
	name string;
	age float64;
}

func Makestruct (age float64) Details {
	mydetisl := Details {
		name: "Ankan",
		age: age,
	}

	return  mydetisl
}