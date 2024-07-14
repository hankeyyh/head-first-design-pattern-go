package main

func main() {
	mallard := NewMallarDuck()
	rubberDuck := NewRubberDuck()
	decoy := NewDecoyDuck()
	modelDuck := NewModelDuck()
	redHeadDuck := NewRedHeadDuck()
	
	// quack
	mallard.performQuack()
	rubberDuck.performQuack()
	decoy.performQuack()
	modelDuck.performQuack()
	redHeadDuck.performQuack()

	// fly
	mallard.performFly()
	rubberDuck.performFly()
	decoy.performFly()
	modelDuck.performFly()
	redHeadDuck.performFly()
	
	rubberDuck.SetFlyBehavior(NewFlyRocketPowered())
	rubberDuck.performFly()
}