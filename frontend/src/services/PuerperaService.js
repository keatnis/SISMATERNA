import HttpService from "./HttpService";
const PuerperaService = {
    async agregarPuerpera(vehiculo) {
        return await HttpService.post("/RegistrarPuerpera", vehiculo);
    },


};
export default PuerperaService;