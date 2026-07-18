use starknet::ContractAddress;
use snforge_std::{declare, ContractClassTrait, DeclareResultTrait};
//import the auto generated dispatchers
use contract_vault::{IVitaVaultDispatcher, IVitaVaultDispatcherTrait};



//start fuction
fn deploy_contarct(name: ByteArray) -> ContractAddress {
    let contract = declare(name).unwrap().contract_class();
    let (contract_address,_) = contract.deploy(@ArrayTrait::new()).unwrap();
    contract_address
}

#[test]
fn test_verify_biometric_proof() {
    //deploy 
    let contract_address = deploy_contarct("VitaVault");
    //create dispatcher to talke to the test network 
    let dispatcher = IVitaVaultDispatcher { contract_address};

    
    //dummy proof but later to change 
    let success = dispatcher.verify_biometric_proof(123);
   //print 
   println!("Status: {}", dispatcher.get_status());

    //valut return true
    assert(success == true, 'Verification Failed');
    assert(dispatcher.get_status() == true, 'data was not saved');

}