pragma circom 2.0.0;

template BiometricVault() {
    signal input fingerprint[13];
    signal output verification_hash;

    // We need an intermediate signal array to hold the multiplied values
    signal squares[13];
    var sum = 0;

    for (var i = 0; i < 13; i++) {
        // Do the multiplication securely in its own constraint
        squares[i] <== fingerprint[i] * fingerprint[i];
        
        // Add it to our running total (addition is perfectly fine!)
        sum = sum + squares[i];
    }
    
    // Output the final sum
    verification_hash <== sum;
}

component main = BiometricVault();