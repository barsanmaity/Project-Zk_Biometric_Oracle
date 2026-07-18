#[starknet::interface]
pub trait IVitaVault<TContractState> {
    fn verify_biometric_proof(ref self: TContractState, proof: felt252) -> bool;
    fn get_status(self :@TContractState) -> bool;
}

#[starknet::contract]
pub mod VitaVault {
    use starknet::storage::StoragePointerReadAccess;
    use starknet::storage::StoragePointerWriteAccess;

    #[storage]
    struct Storage {
        is_verified: bool,
    }
    
    #[event]
    #[derive(Drop, starknet::Event)]
    pub enum Event {
        ProofReceived: ProofReceivedEvent
    }

    #[derive(Drop, starknet::Event)]
    pub struct ProofReceivedEvent {
       pub proof_value : felt252 ,
    }
    

    #[abi(embed_v0)]
    impl VitaVaultImpl of super::IVitaVault<ContractState> {
        fn verify_biometric_proof(ref self: ContractState, proof: felt252) -> bool {

            self.is_verified.write(true);
            self.emit(Event::ProofReceived(ProofReceivedEvent{proof_value: proof}));
            true
        }
        fn get_status(self: @ContractState) -> bool {
                self.is_verified.read()
            }
    }
}