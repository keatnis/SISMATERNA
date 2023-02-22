import HttpService from "./HttpService";
const PuerperaService = {
    async agregarPuerpera(vehiculo) {
        return await HttpService.post("/RegistrarPuerpera", vehiculo);
    },
    async obtenerPuerpera() {
        return await HttpService.get("/ListaPuerpera");
    },

};
export default PuerperaService;