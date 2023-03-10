import HttpService from "./HttpService";
const PuerperaService = {
    async insertPuerpera(puerpera) {
        return await HttpService.post("/RegistrarPuerpera", puerpera);
    },
    async obtenerPuerpera() {
        return await HttpService.get("/ListaPuerpera");
    },
    async login(){
        return await HttpService.post("/auth/login");
    }

};
export default PuerperaService;