package main

type Ray struct {
	O V3
	D V3
}

func CreateRay(O V3, Dun V3) Ray {
	return Ray{O, Dun.Normalize()}
}
