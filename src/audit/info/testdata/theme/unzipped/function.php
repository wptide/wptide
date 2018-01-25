<?php
// Hello class.
class Hello {

	// Private property.
	private $addressee;

    // Constructor.
	public function __construct($addressee = "World") {
		$this-addressee = $addressee;
	}

	// Greeter.
	public function greet() {
		echo "Hello " . $this->addressee;
	}

}

// New Hello.
$helloer = new Hello("Mundo");
// Say Hello.
$helloer->greet();