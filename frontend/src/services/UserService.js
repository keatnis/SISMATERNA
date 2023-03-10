import HttpService from "./HttpService";
const UserService={
    async insertUser(){
        return await HttpService.post('AddUrer');
    }
}

export default UserService;