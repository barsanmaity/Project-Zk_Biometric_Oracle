from processor import extract_fingerprint
import sys

if __name__ == "__main__":
    # Ensure the user provides an audio file when running the script
    if len(sys.argv) < 2:
        print("Usage: python main.py <path_to_audio_file>")
        sys.exit(1)
    
    audio_file = sys.argv[1]
    
    try:
        # 1. Extract the raw decimal fingerprint
        raw_fingerprint = extract_fingerprint(audio_file)
        
        # 2. Scale floats to integers and make them positive for ZK compatibility
        zk_ready_array = [int(abs(x) * 1000) for x in raw_fingerprint]
        
        print(f"Successfully processed: {audio_file}")
        print("-" * 30)
        print("ZK-Ready Integer Array (Send this to the Go Notary):")
        print(zk_ready_array)
        
    except Exception as e:
        print(f"Error processing audio: {e}")