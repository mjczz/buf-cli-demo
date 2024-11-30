
lint:
	buf lint

updatedep:
	buf dep update

generate:
	buf generate

cu: cuget cuput cudel

cuput:
	buf curl --schema . --data '{"pet_type": "PET_TYPE_SNAKE", "name": "老师收拾收拾少时诵l"}' \
      http://localhost:8080/pet.v1.PetStoreService/PutPet

cuget:
	buf curl --schema . --data '{"pet_id": "ids"}' http://localhost:8080/pet.v1.PetStoreService/GetPet

cudel:
	buf curl --schema . --data '{"pet_id": "s"}' http://localhost:8080/pet.v1.PetStoreService/DeletePet

cuhttp-put:
	curl \
        --header "Content-Type: application/json" \
        --data '{"pet_type": "PET_TYPE_SNAKE", "name": "老师收拾收拾少时诵l"}' \
      	http://localhost:8080/pet.v1.PetStoreService/PutPet
