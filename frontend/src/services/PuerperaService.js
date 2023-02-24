import HttpService from "./HttpService";
const PuerperaService = {
    async insertPuerpera(puerpera) {
        return await HttpService.post("/RegistrarPuerpera", puerpera);
    },
    async obtenerPuerpera() {
        return await HttpService.get("/ListaPuerpera");
    },

};
export default PuerperaService;