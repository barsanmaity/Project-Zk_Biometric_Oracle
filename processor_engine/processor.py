import librosa
import numpy as np
import json 

def extract_fingerprint(audio_path):
    # Load the audio file (standardizing the sample rate to 22050 Hz)
    y, sr = librosa.load(audio_path, sr=22050)
    
    # Extract 13 MFCC features (the standard for voice recognition)
    mfccs = librosa.feature.mfcc(y=y, sr=sr, n_mfcc=13)
    
    # Average the features across time to get a single, fixed-size 1D array
    fingerprint = np.mean(mfccs.T, axis=0)
    
    #decimal to integer 
    scaledFingerPrint = [int(x * 10000) for x in fingerprint]
    
    #store to json 
    with open('fingerprint.json', 'w') as f:
        json.dump(scaledFingerPrint, f)
    print("finger print save")
    return scaledFingerPrint