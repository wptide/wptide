// Get object.
var obj = obj | {};

// Add greeter.
obj.greeter = function(msg) {
    // Log to console.
    console.log("Hello " + msg);
}

// Greet.
obj.greeter("World");