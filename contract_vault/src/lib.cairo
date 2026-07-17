#[starknet::interface]
pub trait IVitaVault<TContractState> {
    fn verify_biometric_proof(ref self: TContractState, proof: felt252) -> bool;
}

#[starknet::contract]
pub mod VitaVault {
    use starknet::storage::StoragePointerWriteAccess;

    #[storage]
    struct Storage {
        is_verified: bool,
    }

    #[abi(embed_v0)]
    impl VitaVaultImpl of super::IVitaVault<ContractState> {
        fn verify_biometric_proof(ref self: ContractState, proof: felt252) -> bool {

            self.is_verified.write(true);
            true
        }
    }
}